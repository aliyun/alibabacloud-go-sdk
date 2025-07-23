// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertActionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionIdsShrink(v string) *ListAlertActionsShrinkRequest
	GetAlertActionIdsShrink() *string
	SetAlertActionName(v string) *ListAlertActionsShrinkRequest
	GetAlertActionName() *string
	SetPageNumber(v int32) *ListAlertActionsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertActionsShrinkRequest
	GetPageSize() *int32
	SetType(v string) *ListAlertActionsShrinkRequest
	GetType() *string
}

type ListAlertActionsShrinkRequest struct {
	AlertActionIdsShrink *string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty"`
	// example:
	//
	// testName
	AlertActionName *string `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// FC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAlertActionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertActionsShrinkRequest) GetAlertActionIdsShrink() *string {
	return s.AlertActionIdsShrink
}

func (s *ListAlertActionsShrinkRequest) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *ListAlertActionsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertActionsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertActionsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListAlertActionsShrinkRequest) SetAlertActionIdsShrink(v string) *ListAlertActionsShrinkRequest {
	s.AlertActionIdsShrink = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetAlertActionName(v string) *ListAlertActionsShrinkRequest {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetPageNumber(v int32) *ListAlertActionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetPageSize(v int32) *ListAlertActionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetType(v string) *ListAlertActionsShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
