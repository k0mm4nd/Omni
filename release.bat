@echo off 
FOR /F %%i IN ('go env GOPATH') DO @set OLDGOPATH=%%i
rem echo %OLDGOPATH%
set NEWGOPATH=%~dp0
if "%NEWGOPATH:~-1%" == "\"  set NEWGOPATH=%NEWGOPATH:~0,-1%
rem echo %NEWPATH%
set GOPATH=%NEWGOPATH%
rem go env GOPATH

go get github.com/armon/go-socks5
go get github.com/thoj/go-ircevent
cd bin
go build Omni/Omni -ldflags="-H windowsgui" 
cd ..

set GOPATH=%OLDGOPATH%
rem go env GOPATH
