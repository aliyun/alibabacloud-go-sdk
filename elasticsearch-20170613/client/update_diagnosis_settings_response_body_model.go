// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiagnosisSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDiagnosisSettingsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDiagnosisSettingsResponseBody
	GetResult() *bool
}

type UpdateDiagnosisSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: update successfully
	//
	// 	- false: update failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateDiagnosisSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiagnosisSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDiagnosisSettingsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDiagnosisSettingsResponseBody) SetRequestId(v string) *UpdateDiagnosisSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDiagnosisSettingsResponseBody) SetResult(v bool) *UpdateDiagnosisSettingsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDiagnosisSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
