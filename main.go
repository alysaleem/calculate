package main

import (
   "errors"
   "fmt"
   "strings"
)

func calculate(e string) (int, error) {
   e = strings.ReplaceAll(e, " ", "")
   if len(e) == 0 {
      return 0, errors.New("empty string passed")
   }
   result, _, err := evaluateExpression(e, 0)
   if err != nil {
      return 0, err
   }

   return result, nil
}

// evaluateExpression takes the string and the index from where to start calculating and returns the sum, index and err.
func evaluateExpression(e string, index int) (int, int, error) {
   leftNum := 0
   rightNum := 0
   operator := byte('\n') // set it as invalid operator to start with.
   total := 0

   for index < len(e) {
      ch := e[index]

      if ch >= '0' && ch <= '9' {
         total = total*10 + int(ch-'0')
      } else if ch == '(' {
         // Evaluate the expression inside parentheses recursively
         var innerResult int
         var err error
         innerResult, index, err = evaluateExpression(e, index+1)
         if err != nil {
            return 0, 0, err
         }
         total = innerResult
      } else if ch == ')' {
         // Return the current sub-expression's result and the updated index
         result, err := applyOperator(operator, leftNum, rightNum+total)
         if err != nil {
            return 0, 0, err
         }
         return result, index, nil
      } else if ch == '+' || ch == '-' {
         if operator == byte('\n') {
            operator = '+'
         }
         var err error
         leftNum, err = applyOperator(operator, leftNum, rightNum+total)
         if err != nil {
            return 0, 0, err
         }
         rightNum = 0
         operator = ch
         total = 0
      }

      index++
   }
   result, err := applyOperator(operator, leftNum, rightNum+total)
   if err != nil {
      return 0, 0, err
   }
   // Return the final result and the updated index
   return result, index, nil
}

func applyOperator(operator byte, left, right int) (int, error) {
   switch operator {
   case '+':
      return left + right, nil
   case '-':
      return left - right, nil
   }
   return 0, errors.New("invalid operator")
}

func main() {
   // Test cases
   fmt.Println(calculate("1 + 3")) // Output: -4 <nil>

   fmt.Println(calculate("10  (2 - ( 2 - 3 ))")) // Output: 0 <nil>

   fmt.Println(calculate("10+1+1+1-(10+1+1+1)")) // Output: 0 <nil>

}
