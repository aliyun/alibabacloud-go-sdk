// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachScriptPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAICoachScriptPageRequest
	GetName() *string
	SetPageNumber(v int32) *ListAICoachScriptPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAICoachScriptPageRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListAICoachScriptPageRequest
	GetStatus() *int32
	SetType(v int32) *ListAICoachScriptPageRequest
	GetType() *int32
}

type ListAICoachScriptPageRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAICoachScriptPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageRequest) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageRequest) GetName() *string {
	return s.Name
}

func (s *ListAICoachScriptPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAICoachScriptPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAICoachScriptPageRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAICoachScriptPageRequest) GetType() *int32 {
	return s.Type
}

func (s *ListAICoachScriptPageRequest) SetName(v string) *ListAICoachScriptPageRequest {
	s.Name = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetPageNumber(v int32) *ListAICoachScriptPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetPageSize(v int32) *ListAICoachScriptPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetStatus(v int32) *ListAICoachScriptPageRequest {
	s.Status = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetType(v int32) *ListAICoachScriptPageRequest {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageRequest) Validate() error {
	return dara.Validate(s)
}
