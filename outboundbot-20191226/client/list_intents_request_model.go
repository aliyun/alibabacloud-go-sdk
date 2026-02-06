// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListIntentsRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListIntentsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListIntentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIntentsRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListIntentsRequest
	GetScriptId() *string
}

type ListIntentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ""
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListIntentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntentsRequest) GoString() string {
	return s.String()
}

func (s *ListIntentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListIntentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIntentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListIntentsRequest) SetInstanceId(v string) *ListIntentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntentsRequest) SetKeyword(v string) *ListIntentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListIntentsRequest) SetPageNumber(v int32) *ListIntentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntentsRequest) SetPageSize(v int32) *ListIntentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntentsRequest) SetScriptId(v string) *ListIntentsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListIntentsRequest) Validate() error {
	return dara.Validate(s)
}
