'''
    Given a string s, return true if it is a palindrome, or false otherwise. 
    Rememeber that s only contains alphanumeric characters.
'''

def isPalindrome(s):

    # Convert s to lowercase and keep only aplphanumeric characters
    cleaned = ""

    for char in s:
        if char.isalnum():
            cleaned = cleaned + char.lower()

    # Initialize two pointers
    left = 0
    right = len(cleaned) - 1

    # Compare characters both sides of the string
    while left < right:
        if cleaned[left] != cleaned[right]:
            return False

        left = left + 1
        right = right - 1

    # Return true if no mismatches found
    return True

print(isPalindrome("A man, a plan, a canal: Panama"))
print(isPalindrome("race a car"))
print(isPalindrome(" "))
