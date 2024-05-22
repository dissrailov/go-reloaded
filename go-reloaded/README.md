# GO-RELOADED

This project is a text modification tool written in Go that adheres to good coding practices. The tool takes as input a text file requiring modifications and outputs the modified text to another file. The following list outlines the modifications the program executes:

                             ***HOW TO USE***
    1. Clone the repository
    2. Use go run . sample.txt result.txt



1. **Hexadecimal Conversion:**
   - Replace every instance of `(hex)` with the decimal version of the preceding word (hexadecimal number).
   - Example: `"1E (hex) files were added" -> "30 files were added"`

2. **Binary Conversion:**
   - Replace every instance of `(bin)` with the decimal version of the preceding word (binary number).
   - Example: `"It has been 10 (bin) years" -> "It has been 2 years"`

3. **Uppercase Conversion:**
   - Replace every instance of `(up)` with the uppercase version of the preceding word.
   - Example: `"Ready, set, go (up)!" -> "Ready, set, GO!"`

4. **Lowercase Conversion:**
   - Replace every instance of `(low)` with the lowercase version of the preceding word.
   - Example: `"I should stop SHOUTING (low)" -> "I should stop shouting"`

5. **Capitalized Conversion:**
   - Replace every instance of `(cap)` with the capitalized version of the preceding word.
   - Example: `"Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"`

6. **Customized Conversion:**
   - For `(low)`, `(up)`, and `(cap)`, if followed by a number, convert the specified number of words accordingly.
   - Example: `"This is so exciting (up, 2)" -> "This is SO EXCITING"`

7. **Punctuation Formatting:**
   - Ensure proper spacing around punctuation marks (., ,, !, ?, :, and ;).
   - Handle groups of punctuation like ... or !? by formatting accordingly.
   - Example: `"I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"`
     `"I was thinking ... You were right" -> "I was thinking... You were right"`

8. **A to An Conversion:**
   - Convert every instance of "a" to "an" if the next word begins with a vowel (a, e, i, o, u) or "h".
   - Example: `"There it was. A amazing rock!" -> "There it was. An amazing rock!"`

9. **Quotation Marks Placement:**
   - Place quotation marks (' ') to the right and left of words enclosed between them, without any spaces.
   - Example: `"I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"`

10. **Multiple Word Quotation Marks Placement:**
    - If there are more than one word between the two ' ' marks, place the marks next to the corresponding words.
    - Example: `"As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"`

                         **EXCEPTION**
                1. "don't (command)" don't work 
                2. there is one space everywhere


The program ensures readability, adherence to coding standards, and includes unit test files for thorough testing.