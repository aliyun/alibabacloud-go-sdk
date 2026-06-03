// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptPublishHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptPublishHistoriesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListScriptPublishHistoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptPublishHistoriesRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListScriptPublishHistoriesRequest
	GetScriptId() *string
}

type ListScriptPublishHistoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
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
	// b0f35dd1-0337-402e-9c4f-3a6c2426950a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListScriptPublishHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptPublishHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListScriptPublishHistoriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptPublishHistoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptPublishHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptPublishHistoriesRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptPublishHistoriesRequest) SetInstanceId(v string) *ListScriptPublishHistoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptPublishHistoriesRequest) SetPageNumber(v int32) *ListScriptPublishHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptPublishHistoriesRequest) SetPageSize(v int32) *ListScriptPublishHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptPublishHistoriesRequest) SetScriptId(v string) *ListScriptPublishHistoriesRequest {
	s.ScriptId = &v
	return s
}

func (s *ListScriptPublishHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
