// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenDiagnosisResponseBody
	GetRequestId() *string
	SetResult(v bool) *OpenDiagnosisResponseBody
	GetResult() *bool
}

type OpenDiagnosisResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s OpenDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDiagnosisResponseBody) GetResult() *bool {
	return s.Result
}

func (s *OpenDiagnosisResponseBody) SetRequestId(v string) *OpenDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDiagnosisResponseBody) SetResult(v bool) *OpenDiagnosisResponseBody {
	s.Result = &v
	return s
}

func (s *OpenDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
