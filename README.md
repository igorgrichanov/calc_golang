Hi!

This repository contains code of a calculator for Arabic and Roman numbers written in Golang.



The calculator can perform addition, subtraction, multiplication and division operations with two numbers: a + b, a - b, a * b, a / b.

Numbers to calculate is transferred in one line. Solutions in which each number and arithmetic operation is passed from a new line are considered incorrect.

The calculator can work with both Arabic (1,2,3,4,5…) and Roman (I,II,III,IV,V…) numbers.

The calculator accepts input numbers from 1 to 10 inclusive, no more. At the output, the numbers are not limited in size and can be any.
The calculator can only work with integers.

The calculator can only work with Arabic or Roman numerals at the same time, when the user enters a string like 3 + II, the calculator indicates an error and stop working.

When entering Roman numerals, the answer will be displayed in Roman numerals, respectively, when entering Arabic - the answer will be in Arabic.

If the user enters inappropriate numbers, the application displays an error in the terminal and exits.

If the user enters a string that does not match one of the above arithmetic operations, the application displays an error and exits.

The result of the division operation is an integer, the remainder is discarded.

The result of the calculator with Arabic numbers can be negative numbers and zero. The result of the calculator with Roman numerals can only be positive numbers, if the result of the work is less than one, the program should indicate an exception.
