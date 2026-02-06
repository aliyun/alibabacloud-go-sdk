// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalQuestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListGlobalQuestionsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListGlobalQuestionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGlobalQuestionsRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListGlobalQuestionsRequest
	GetScriptId() *string
}

type ListGlobalQuestionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListGlobalQuestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalQuestionsRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalQuestionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGlobalQuestionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGlobalQuestionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGlobalQuestionsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListGlobalQuestionsRequest) SetInstanceId(v string) *ListGlobalQuestionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGlobalQuestionsRequest) SetPageNumber(v int32) *ListGlobalQuestionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGlobalQuestionsRequest) SetPageSize(v int32) *ListGlobalQuestionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGlobalQuestionsRequest) SetScriptId(v string) *ListGlobalQuestionsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListGlobalQuestionsRequest) Validate() error {
	return dara.Validate(s)
}
