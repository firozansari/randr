## randr - Rename And Replace

randr is a basic open source command line tool to recursively search for source text across all files in the provided folder, and 
replace it with the target text. With that, randr will also recursively search for files or folders name which contains source text 
and rename it with the text target text. 

Replacing content as well as renaming files/folders with a single command makes randr a very handy tool to refactor your code. With 
that, you can also specify in the tool if you want case sensitive or case insensitive search for renaming or replacing text.

To illustrate the usage of randr, consider you have a project folder with the following structure:

``` csharp
// Project Folder: c:\projects\payroll\

c:\projects\payroll\Models\Worker.cs
public class Worker {}

c:\projects\payroll\Services\WorkerService.cs
public class WorkerService {}

c:\projects\payroll\Repositories\WorkerRepository.cs
public class WorkerRepository {}

c:\projects\payroll\Helpers\Worker\Extensions.cs
public static class Extensions {}

```

Let's refactor above project and change `Worker` entities to `Employee`. `randr` Command will look like the following:

```
randr.exe -find=Worker -replace=Employee -match=true -location=c:\projects\payroll\
```

After running randr tool successfully, our example project folder will look like following:

``` csharp
// Project Folder: c:\projects\payroll\

c:\projects\payroll\Models\Employee.cs
public class Employee {}

c:\projects\payroll\Services\EmployeeService.cs
public class EmployeeService {}

c:\projects\payroll\Repositories\EmployeeRepository.cs
public class EmployeeRepository {}

c:\projects\payroll\Helpers\Employee\Extensions.cs
public static class Extensions {}

```

## Usage
```
Usage:
  rand -find=source -replace=target [-match=true/false] -location=directory

Arguments:
  -f, -find       find text to rename or replace
  -r, -replace    replace text to
  -m, -match      match case -- true: case sensitive (default); false: case insensitive
  -l, -location   directory location to process files and sub directories

  -h, -?, -help   show command usage and examples

Examples:
  randr.exe -find=worker -replace=employee -match=true -location=c:\projectss\payroll\
  randr.exe -find=Notes -replace=contents -match=false -location=C:\documents\daily\
```  

## Build randr Source
To build and create executable file, goto randr project source folder and run following command:

``` terminal
go build
```
Go compiler will generate `randr.exe` in the source folder.

## License
The MIT License

Copyright (c) 2019 Firoz Ansari. https://firozansari.com

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.