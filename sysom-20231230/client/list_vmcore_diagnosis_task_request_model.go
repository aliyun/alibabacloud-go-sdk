// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVmcoreDiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int64) *ListVmcoreDiagnosisTaskRequest
	GetDays() *int64
}

type ListVmcoreDiagnosisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	Days *int64 `json:"days,omitempty" xml:"days,omitempty"`
}

func (s ListVmcoreDiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVmcoreDiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *ListVmcoreDiagnosisTaskRequest) GetDays() *int64 {
	return s.Days
}

func (s *ListVmcoreDiagnosisTaskRequest) SetDays(v int64) *ListVmcoreDiagnosisTaskRequest {
	s.Days = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
