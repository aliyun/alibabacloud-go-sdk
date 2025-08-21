// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseDiagnosisResponseBody
	GetRequestId() *string
	SetResult(v bool) *CloseDiagnosisResponseBody
	GetResult() *bool
}

type CloseDiagnosisResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseDiagnosisResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CloseDiagnosisResponseBody) SetRequestId(v string) *CloseDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDiagnosisResponseBody) SetResult(v bool) *CloseDiagnosisResponseBody {
	s.Result = &v
	return s
}

func (s *CloseDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
