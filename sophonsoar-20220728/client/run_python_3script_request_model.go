// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPython3ScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeName(v string) *RunPython3ScriptRequest
	GetNodeName() *string
	SetParams(v string) *RunPython3ScriptRequest
	GetParams() *string
	SetPlaybookUuid(v string) *RunPython3ScriptRequest
	GetPlaybookUuid() *string
	SetPythonScript(v string) *RunPython3ScriptRequest
	GetPythonScript() *string
	SetPythonVersion(v string) *RunPython3ScriptRequest
	GetPythonVersion() *string
}

type RunPython3ScriptRequest struct {
	// The name of the node in the playbook.
	//
	// example:
	//
	// python3_3
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The input parameters of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the UUIDs of playbooks.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The Python3 script.
	//
	// example:
	//
	// import logging
	//
	// def execute (params):
	//
	//   #ip = params[\\"ip\\"]
	//
	//   #logging.info("enter execute,ip is "+ip)
	//
	//   success=True
	//
	//   message=\\"OK\\"
	//
	//   data=[]
	//
	//   return (success,message,data)
	PythonScript *string `json:"PythonScript,omitempty" xml:"PythonScript,omitempty"`
	// example:
	//
	// python2.0
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
}

func (s RunPython3ScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPython3ScriptRequest) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *RunPython3ScriptRequest) GetParams() *string {
	return s.Params
}

func (s *RunPython3ScriptRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *RunPython3ScriptRequest) GetPythonScript() *string {
	return s.PythonScript
}

func (s *RunPython3ScriptRequest) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *RunPython3ScriptRequest) SetNodeName(v string) *RunPython3ScriptRequest {
	s.NodeName = &v
	return s
}

func (s *RunPython3ScriptRequest) SetParams(v string) *RunPython3ScriptRequest {
	s.Params = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPlaybookUuid(v string) *RunPython3ScriptRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPythonScript(v string) *RunPython3ScriptRequest {
	s.PythonScript = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPythonVersion(v string) *RunPython3ScriptRequest {
	s.PythonVersion = &v
	return s
}

func (s *RunPython3ScriptRequest) Validate() error {
	return dara.Validate(s)
}
