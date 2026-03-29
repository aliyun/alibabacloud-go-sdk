// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDataDiagnosesShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDataDiagnosesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataDiagnosesShrinkRequest
	GetPageSize() *int32
	SetTypesShrink(v string) *ListDataDiagnosesShrinkRequest
	GetTypesShrink() *string
}

type ListDataDiagnosesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDataDiagnosesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataDiagnosesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataDiagnosesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataDiagnosesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListDataDiagnosesShrinkRequest) SetInstanceId(v string) *ListDataDiagnosesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataDiagnosesShrinkRequest) SetPageNumber(v int32) *ListDataDiagnosesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataDiagnosesShrinkRequest) SetPageSize(v int32) *ListDataDiagnosesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataDiagnosesShrinkRequest) SetTypesShrink(v string) *ListDataDiagnosesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListDataDiagnosesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
