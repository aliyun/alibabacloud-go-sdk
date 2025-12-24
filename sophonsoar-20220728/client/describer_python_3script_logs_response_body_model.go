// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescriberPython3ScriptLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescriberPython3ScriptLogsResponseBody
	GetRequestId() *string
	SetRunResult(v string) *DescriberPython3ScriptLogsResponseBody
	GetRunResult() *string
}

type DescriberPython3ScriptLogsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D22D8A0C-6E86-57B2-A142-929184122AB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The operational logs of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "logs": [
	//
	//         {
	//
	//             "message": "function input is {}"
	//
	//         }
	//
	//     ]
	//
	// }
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s DescriberPython3ScriptLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescriberPython3ScriptLogsResponseBody) GetRunResult() *string {
	return s.RunResult
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRequestId(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRunResult(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RunResult = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
