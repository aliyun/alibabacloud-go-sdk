// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDataDiagnosisJobsShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListDataDiagnosisJobsShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataDiagnosisJobsShrinkRequest
	GetPageSize() *string
	SetStatus(v string) *ListDataDiagnosisJobsShrinkRequest
	GetStatus() *string
	SetTypesShrink(v string) *ListDataDiagnosisJobsShrinkRequest
	GetTypesShrink() *string
}

type ListDataDiagnosisJobsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Initializing
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDataDiagnosisJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisJobsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataDiagnosisJobsShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataDiagnosisJobsShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataDiagnosisJobsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDataDiagnosisJobsShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListDataDiagnosisJobsShrinkRequest) SetInstanceId(v string) *ListDataDiagnosisJobsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataDiagnosisJobsShrinkRequest) SetPageNumber(v string) *ListDataDiagnosisJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataDiagnosisJobsShrinkRequest) SetPageSize(v string) *ListDataDiagnosisJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataDiagnosisJobsShrinkRequest) SetStatus(v string) *ListDataDiagnosisJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDataDiagnosisJobsShrinkRequest) SetTypesShrink(v string) *ListDataDiagnosisJobsShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListDataDiagnosisJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
