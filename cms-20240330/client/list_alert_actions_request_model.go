// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionIds(v []*string) *ListAlertActionsRequest
	GetAlertActionIds() []*string
	SetAlertActionName(v string) *ListAlertActionsRequest
	GetAlertActionName() *string
	SetPageNumber(v int32) *ListAlertActionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertActionsRequest
	GetPageSize() *int32
	SetType(v string) *ListAlertActionsRequest
	GetType() *string
}

type ListAlertActionsRequest struct {
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
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

func (s ListAlertActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertActionsRequest) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *ListAlertActionsRequest) GetAlertActionName() *string {
	return s.AlertActionName
}

func (s *ListAlertActionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertActionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertActionsRequest) GetType() *string {
	return s.Type
}

func (s *ListAlertActionsRequest) SetAlertActionIds(v []*string) *ListAlertActionsRequest {
	s.AlertActionIds = v
	return s
}

func (s *ListAlertActionsRequest) SetAlertActionName(v string) *ListAlertActionsRequest {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsRequest) SetPageNumber(v int32) *ListAlertActionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsRequest) SetPageSize(v int32) *ListAlertActionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsRequest) SetType(v string) *ListAlertActionsRequest {
	s.Type = &v
	return s
}

func (s *ListAlertActionsRequest) Validate() error {
	return dara.Validate(s)
}
