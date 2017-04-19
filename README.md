# homographs
Homographs: brutefind homographs within a font


This tool will generate a table of all homographs (identical characters) within a font and create a list of all possible domains.

[https://isc.sans.edu/forums/diary/Tool+to+Detect+Active+Phishing+Attacks+Using+Unicode+LookAlike+Domains/22310/](https://isc.sans.edu/forums/diary/Tool+to+Detect+Active+Phishing+Attacks+Using+Unicode+LookAlike+Domains/22310/)

# Usage

```
$ homographs --font-file ./Arial.ttf apple.com
Homographs: brutefind homographs within a font
by Remco Verhoef (@remco_verhoef) of DutchSec
---------------------------------------------------
[+] Homograph table
A(0x41) | Α(0x391), А(0x410)
B(0x42) | Β(0x392), В(0x412)
C(0x43) | Ϲ(0x3f9), С(0x421)
E(0x45) | Ε(0x395), Е(0x415)
H(0x48) | Η(0x397), Н(0x41d)
I(0x49) | Ι(0x399), І(0x406), Ӏ(0x4c0)
J(0x4a) | Ј(0x408)
K(0x4b) | Κ(0x39a)
M(0x4d) | Μ(0x39c), М(0x41c)
N(0x4e) | Ν(0x39d)
O(0x4f) | Ο(0x39f), О(0x41e)
P(0x50) | Ρ(0x3a1), Р(0x420)
S(0x53) | Ѕ(0x405)
T(0x54) | Τ(0x3a4), Т(0x422)
X(0x58) | Χ(0x3a7), Х(0x425)
Y(0x59) | Υ(0x3a5)
Z(0x5a) | Ζ(0x396)
a(0x61) | а(0x430)
c(0x63) | ϲ(0x3f2), с(0x441)
d(0x64) | ԁ(0x501)
e(0x65) | е(0x435)
g(0x67) | ɡ(0x261)
h(0x68) | һ(0x4bb)
i(0x69) | і(0x456)
j(0x6a) | ϳ(0x3f3), ј(0x458)
l(0x6c) | Ɩ(0x196), ӏ(0x4cf)
o(0x6f) | ο(0x3bf), о(0x43e)
p(0x70) | р(0x440)
s(0x73) | ѕ(0x455)
v(0x76) | ν(0x3bd)
x(0x78) | х(0x445)
y(0x79) | у(0x443)

[+] Homograph domains
аpple.com xn--pple-43d.com
арple.com xn--ple-5cd4f.com
аpрle.com xn--ple-5cd5f.com
аppƖe.com xn--ppe-35a144a.com
аppӏe.com xn--ppe-5cd11f.com
аpplе.com xn--ppl-5cd2a.com
аррle.com xn--le-6kc8da.com
арpƖe.com xn--pe-4xa67xkc.com
арpӏe.com xn--pe-6kc8d62b.com
арplе.com xn--pl-6kcw6c.com
аpрƖe.com xn--pe-4xa67xlc.com
аpрӏe.com xn--pe-6kc9d52b.com
аpрlе.com xn--pl-6kcw7c.com
аppƖе.com xn--pp-5xa57x2a.com
аppӏе.com xn--pp-6kcw38e.com
аррƖe.com xn--e-5pa11r3ba.com
аррӏe.com xn--e-7sb2ca92e.com
аррlе.com xn--l-7sbq6ba.com
арpƖе.com xn--p-6pa01rwa6c.com
арpӏе.com xn--p-7sbq6b43c.com
аpрƖе.com xn--p-6pa01rwa7c.com
аpрӏе.com xn--p-7sbq7b33c.com
аррƖе.com xn--7ha54kqa6ba.com
аррӏе.com xn--80ak6aa92e.com
aрple.com xn--aple-g6d.com
aррle.com xn--ale-0eda.com
aрpƖe.com xn--ape-35a225a.com
aрpӏe.com xn--ape-0ed03e.com
aрplе.com xn--apl-tdd6c.com
aррƖe.com xn--ae-4xa14ya.com
aррӏe.com xn--ae-lmca92e.com
aррlе.com xn--al-olc6ba.com
aрpƖе.com xn--ap-5xa89xrb.com
aрpӏе.com xn--ap-olc6b43c.com
aррƖе.com xn--a-6pa72rhba.com
aррӏе.com xn--a-jtb6aa92e.com
apрle.com xn--aple-h6d.com
apрƖe.com xn--ape-35a325a.com
apрӏe.com xn--ape-1ed92e.com
apрlе.com xn--apl-tdd7c.com
apрƖе.com xn--ap-5xa89xsb.com
apрӏе.com xn--ap-olc7b33c.com
appƖe.com xn--appe-2eb.com
appƖе.com xn--app-45a964a.com
appӏe.com xn--appe-xre.com
appӏе.com xn--app-tdd38e.com
applе.com xn--appl-y4d.com
``` 
