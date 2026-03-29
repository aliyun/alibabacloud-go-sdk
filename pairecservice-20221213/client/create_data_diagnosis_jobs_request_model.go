// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiagnosisId(v string) *CreateDataDiagnosisJobsRequest
	GetDataDiagnosisId() *string
	SetEndDate(v string) *CreateDataDiagnosisJobsRequest
	GetEndDate() *string
	SetInstanceId(v string) *CreateDataDiagnosisJobsRequest
	GetInstanceId() *string
	SetStartDate(v string) *CreateDataDiagnosisJobsRequest
	GetStartDate() *string
}

type CreateDataDiagnosisJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	DataDiagnosisId *string `json:"DataDiagnosisId,omitempty" xml:"DataDiagnosisId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CreateDataDiagnosisJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisJobsRequest) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisJobsRequest) GetDataDiagnosisId() *string {
	return s.DataDiagnosisId
}

func (s *CreateDataDiagnosisJobsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateDataDiagnosisJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDataDiagnosisJobsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateDataDiagnosisJobsRequest) SetDataDiagnosisId(v string) *CreateDataDiagnosisJobsRequest {
	s.DataDiagnosisId = &v
	return s
}

func (s *CreateDataDiagnosisJobsRequest) SetEndDate(v string) *CreateDataDiagnosisJobsRequest {
	s.EndDate = &v
	return s
}

func (s *CreateDataDiagnosisJobsRequest) SetInstanceId(v string) *CreateDataDiagnosisJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDataDiagnosisJobsRequest) SetStartDate(v string) *CreateDataDiagnosisJobsRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDataDiagnosisJobsRequest) Validate() error {
	return dara.Validate(s)
}
