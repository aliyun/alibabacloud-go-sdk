// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *ListKeyVersionsRequest
	GetKeyId() *string
	SetPageNumber(v int32) *ListKeyVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKeyVersionsRequest
	GetPageSize() *int32
}

type ListKeyVersionsRequest struct {
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 101.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKeyVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ListKeyVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKeyVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeyVersionsRequest) SetKeyId(v string) *ListKeyVersionsRequest {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageNumber(v int32) *ListKeyVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageSize(v int32) *ListKeyVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeyVersionsRequest) Validate() error {
	return dara.Validate(s)
}
