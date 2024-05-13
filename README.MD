made with lav!

Unfortunately test is not working.
What i tried ?

I tried to match the bash output and the wanted one;
I copy/paste the desired output in to control file, because of tearing of letters, it did not work.

I tried to make /test folder. Did not work (problem was lying there ^^)

I learned to `cat` is not a good stuff for this kind of job, switch to `tee` its echoing/writing the file
its all good its alright but the problem is; the last `$` symbol has to be in the 8. line without space
but nope! Its either 7 `$` or all lined up.

Usage is declared and checking/handling the errors inside of program.

Its easy: 

If test file will be in your directory:
go run main.go "aWord" | cat -e

If u want to remove the test file for avoiding -possibly- conflicting stuff:
go run . "aWord" | cat -e

Tried/tested with :

https://talent.uniworkhub.com/intra/uniworkhub/div-01/ascii-art?event=63
