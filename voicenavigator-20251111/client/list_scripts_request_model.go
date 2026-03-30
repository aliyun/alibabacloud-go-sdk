// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptsRequest
	GetInstanceId() *string
	SetName(v string) *ListScriptsRequest
	GetName() *string
	SetNumber(v string) *ListScriptsRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListScriptsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptsRequest
	GetPageSize() *int32
	SetScriptIds(v []*string) *ListScriptsRequest
	GetScriptIds() []*string
	SetStatus(v string) *ListScriptsRequest
	GetStatus() *string
}

type ListScriptsRequest struct {
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
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScriptIds []*string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty" type:"Repeated"`
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsRequest) GetName() *string {
	return s.Name
}

func (s *ListScriptsRequest) GetNumber() *string {
	return s.Number
}

func (s *ListScriptsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsRequest) GetScriptIds() []*string {
	return s.ScriptIds
}

func (s *ListScriptsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListScriptsRequest) SetInstanceId(v string) *ListScriptsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsRequest) SetName(v string) *ListScriptsRequest {
	s.Name = &v
	return s
}

func (s *ListScriptsRequest) SetNumber(v string) *ListScriptsRequest {
	s.Number = &v
	return s
}

func (s *ListScriptsRequest) SetPageNumber(v int32) *ListScriptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsRequest) SetPageSize(v int32) *ListScriptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptsRequest) SetScriptIds(v []*string) *ListScriptsRequest {
	s.ScriptIds = v
	return s
}

func (s *ListScriptsRequest) SetStatus(v string) *ListScriptsRequest {
	s.Status = &v
	return s
}

func (s *ListScriptsRequest) Validate() error {
	return dara.Validate(s)
}
