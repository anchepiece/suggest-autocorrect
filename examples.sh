#!/bin/sh

echo == Exact match
./suggest-autocorrect -q "test" -c "test, blank, fgrep"
echo 

echo == Autocorrected 
./suggest-autocorrect -q "gferp" -c "tests, blank, fgrep"
echo 

echo == Autocorrected 
./suggest-autocorrect -q "kittnes" -c "ribbons, bitten, sitting, kittens"
echo 

echo == Autocorrect disabled
./suggest-autocorrect -q "kittnes" -c "ribbons, bitten, sitting, kittens" -d
echo 


echo == Few good suggestions
./suggest-autocorrect -q "kittens" -c "tests, mittens, sitting"
echo 

echo == One good suggestion
./suggest-autocorrect -q "initt" -c "init, sine"
echo 

echo == No good suggestion
./suggest-autocorrect -q "unique" -c "tests, blank, fgrep"
echo 

echo == No good suggestion
./suggest-autocorrect -q "Arnold schwartzenegger" -c "Arnold Schwarzenegger, blank, fgrep"
echo 


