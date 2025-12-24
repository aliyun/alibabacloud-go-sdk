// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPython3ScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunPython3ScriptResponseBody
	GetRequestId() *string
	SetRunResult(v string) *RunPython3ScriptResponseBody
	GetRunResult() *string
}

type RunPython3ScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F210521C-D9BF-5264-8369-83EDDC617DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "requestUuid": "fe240b98-27b1-4a36-aec1-550b894318d9",
	//
	//     "content": {
	//
	//         "resultData": [],
	//
	//         "success": true
	//
	//     }
	//
	// }
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s RunPython3ScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPython3ScriptResponseBody) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPython3ScriptResponseBody) GetRunResult() *string {
	return s.RunResult
}

func (s *RunPython3ScriptResponseBody) SetRequestId(v string) *RunPython3ScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPython3ScriptResponseBody) SetRunResult(v string) *RunPython3ScriptResponseBody {
	s.RunResult = &v
	return s
}

func (s *RunPython3ScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
