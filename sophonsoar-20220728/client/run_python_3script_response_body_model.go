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
	// The ID of the request. Alibaba Cloud generates this unique ID for the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// F210521C-D9BF-5264-8369-83EDDC617DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned after the script is run.
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
