// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtFailureTime(v string) *CreateDiagnosisRequest
	GetGmtFailureTime() *string
	SetInstanceId(v string) *CreateDiagnosisRequest
	GetInstanceId() *string
	SetProblemCategory(v string) *CreateDiagnosisRequest
	GetProblemCategory() *string
}

type CreateDiagnosisRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtFailureTime *string `json:"GmtFailureTime,omitempty" xml:"GmtFailureTime,omitempty"`
	// example:
	//
	// dsw-5bk19******n97w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// InstanceStartFailed
	ProblemCategory *string `json:"ProblemCategory,omitempty" xml:"ProblemCategory,omitempty"`
}

func (s CreateDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisRequest) GetGmtFailureTime() *string {
	return s.GmtFailureTime
}

func (s *CreateDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDiagnosisRequest) GetProblemCategory() *string {
	return s.ProblemCategory
}

func (s *CreateDiagnosisRequest) SetGmtFailureTime(v string) *CreateDiagnosisRequest {
	s.GmtFailureTime = &v
	return s
}

func (s *CreateDiagnosisRequest) SetInstanceId(v string) *CreateDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDiagnosisRequest) SetProblemCategory(v string) *CreateDiagnosisRequest {
	s.ProblemCategory = &v
	return s
}

func (s *CreateDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
