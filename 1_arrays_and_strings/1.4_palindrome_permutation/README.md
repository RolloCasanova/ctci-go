# Palindrome Permutation

Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.

## Example

``` nocode
 Input: Tact Coa
Output: True (permutations: "taco cat", "atco cta", etc.)
```

## Hints

<details>
    <summary>#106</summary>
    You do not have to (and should not) generate all permutations. This would be very inefficient.
</details>

<details>
    <summary>#121</summary>
    What characteristics would a string that is a permutation of a palindrome have?
</details>

<details>
    <summary>#134</summary>
    Have you tried a hash table? You should be able to get this down to O(N) time.
</details>

<details>
    <summary>#136</summary>
    Can you reduce the space usage by using a bit vector?
</details>
