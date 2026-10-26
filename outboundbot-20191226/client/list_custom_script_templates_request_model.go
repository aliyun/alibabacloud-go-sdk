// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomScriptTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCustomScriptTemplatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCustomScriptTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomScriptTemplatesRequest
	GetPageSize() *int32
}

type ListCustomScriptTemplatesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 0ec0c897-b92c-40e4-9ad7-e6e4f5ce13bb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for paging. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for paging. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCustomScriptTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomScriptTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomScriptTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomScriptTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomScriptTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomScriptTemplatesRequest) SetInstanceId(v string) *ListCustomScriptTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomScriptTemplatesRequest) SetPageNumber(v int32) *ListCustomScriptTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomScriptTemplatesRequest) SetPageSize(v int32) *ListCustomScriptTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomScriptTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
