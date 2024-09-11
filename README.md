# calc_golang

Implementation of the calculator that operates with Arabic and Roman numbers

## Description

The calculator can perform 4 operations with two numbers:
- Addition: a + b;
- Subtraction: a - b;
- Multiplication: a * b;
- Division: a / b.

## Important notes

- Numbers to calculate is transferred in one line;
- The calculator can only work with Arabic (1,2,3,4,5…) or Roman (I,II,III,IV,V…) numbers at the same time. There cannot be an Arabic and a Roman number in one input line;
- The calculator accepts input numbers from 1 to 10 inclusive. At the output, the numbers are not limited in size and can be any.
- The calculator can only work with integers;
- The input expression can be separated with spaces (4 + 3) or not (4+3).
- When entering Roman numbers, the answer will be displayed in Roman numerals, when entering Arabic - the answer will be in Arabic.
- If the user enters inappropriate numbers, the application displays an error in the terminal and exits.
- If the user enters a string that does not match one of the above arithmetic operations, the application displays an error and exits.
- The result of the division operation is an integer, the remainder is discarded.
- The result of the calculator with Arabic numbers can be negative numbers and zero. The result of the calculator with Roman numerals can only be positive numbers. If the result of the work is less than one, the application displays an error and exits.
