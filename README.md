# serve

Usage: ./serve directory-name

Example:

```
term1
  
  ./serve ~/files  
  👂  @ 10.0.13.136:8000  
  
  downloading  /Users/viresh/files/  
  downloading  /Users/viresh/files/test1  
```
  
  
```
term2

  $ curl http://localhost:8000/
  <pre>
  <a href="test1">test1</a>
  <a href="test2">test2</a>
  <a href="test3">test3</a>
  </pre>


  $ wget http://localhost:8000/test1
  --2017-11-16 14:50:16--  http://localhost:8000/test1
  Resolving localhost... ::1, 127.0.0.1
  Connecting to localhost|::1|:8000... connected.
  HTTP request sent, awaiting response... 200 OK
  Length: 0 [text/plain]
  Saving to: 'test1.1'

  test1.1                 [ <=>              ]       0  --.-KB/s    in 0s

  2017-11-16 14:50:16 (0.00 B/s) - 'test1.1' saved [0/0]
```
