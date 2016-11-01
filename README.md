# go-string-encrypt

This is a really simple, small package, using code swiped from stackoverflow,
which should purely save re-creating a bunch of lines in multiple other
projects.

## Usage

Import it with a more manageable identifier:

```
import (
  ...

  gse "github.com/janeprather/go-string-encrypt"
)
```

Use it:

```
myString := "my text i want encrypted"
myKey := gse.GenerateKey()

var err error

myStringCrypt, err = gse.Encrypt(myKey, myString)

// do something with myStringCrypt, it's an ASCII string

myStringAgain, err := gse.Decrypt(myKey, myStringCrypt)

// myStringAgain should now be the original clear text
```

Note that the _key_ can be any 32byte string represented in base64.  The
GenerateKey() function exists to assist in quickly generating a new random
string suitable for the task.

## A Thought On Security

Note that the key will allow decryption of the data, therefore, if your
encrypted strings are stored somewhere like a database or other shared resource,
you can reduce the risk of that data being compromised by keeping the key
on the local machine(s) that do the encrypting/decrypting instead of on the shared
data resources.  If a database gets stolen, at least it won't contain your
decryption key!

If, however, you are in a situation where you need to keep the key and the
encrypted strings on the same disk/volume/qtree/whatever, then this is more
of an obfuscation than real security. Anybody who can view the source code
for the app can easily see how to use the key to decrypt the encrypted strings.
People without the source code may be able to glean some information about it
from the binary executable with reverse engineering tools.  In either case,
a determined user could decrypt the encrypted strings with access to 1) the
key; 2) the encrypted strings; 3) the source, or possibly even binary of the
app.  It still feels a whole lot better than leaving sensitive data in
cleartext.
