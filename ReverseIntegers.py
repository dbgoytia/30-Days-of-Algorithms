
# Given a 32-bit signed integer, reverse digits of an integer.
# Easy solution, given that to reverse a string you can use a stack and continuously push and pop the elements.

# Author: Diego Goytia

class Solution:
    
    def reverse(self, x: int) -> int:
        stack = []
        res = ""
        x = str(x)
        
        # Save the sign of the number
        if x[0] == '-':
            res = '-'

        # Append and pop operations to reverse the string
        for num in x:
            # Ignore the sign during pushing to stack
            if (num != '-'):
                stack.append(num)
        while len(stack) != 0:
            res = res + stack.pop()
            
        # Check if it's a 32 bit signed integer, else return zero.
        if not (-2**31)<= int(res) <= (2**31-1):
            return  0 
                    
        # Remove all leading zeros from the number:
        return int(res)
    
    
