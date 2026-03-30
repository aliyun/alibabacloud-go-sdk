// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptsShrinkRequest
	GetInstanceId() *string
	SetName(v string) *ListScriptsShrinkRequest
	GetName() *string
	SetNumber(v string) *ListScriptsShrinkRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListScriptsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptsShrinkRequest
	GetPageSize() *int32
	SetScriptIdsShrink(v string) *ListScriptsShrinkRequest
	GetScriptIdsShrink() *string
	SetStatus(v string) *ListScriptsShrinkRequest
	GetStatus() *string
}

type ListScriptsShrinkRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 13816111993
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScriptIdsShrink *string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScriptsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListScriptsShrinkRequest) GetNumber() *string {
	return s.Number
}

func (s *ListScriptsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsShrinkRequest) GetScriptIdsShrink() *string {
	return s.ScriptIdsShrink
}

func (s *ListScriptsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListScriptsShrinkRequest) SetInstanceId(v string) *ListScriptsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetName(v string) *ListScriptsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetNumber(v string) *ListScriptsShrinkRequest {
	s.Number = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetPageNumber(v int32) *ListScriptsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetPageSize(v int32) *ListScriptsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetScriptIdsShrink(v string) *ListScriptsShrinkRequest {
	s.ScriptIdsShrink = &v
	return s
}

func (s *ListScriptsShrinkRequest) SetStatus(v string) *ListScriptsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListScriptsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
