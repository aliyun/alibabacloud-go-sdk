// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesByKeyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *ListAliasesByKeyIdRequest
	GetKeyId() *string
	SetPageNumber(v int32) *ListAliasesByKeyIdRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAliasesByKeyIdRequest
	GetPageSize() *int32
}

type ListAliasesByKeyIdRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The number of the page to return.
	//
	// Valid values: an integer that is greater than 0.
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
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesByKeyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesByKeyIdRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ListAliasesByKeyIdRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAliasesByKeyIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliasesByKeyIdRequest) SetKeyId(v string) *ListAliasesByKeyIdRequest {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageNumber(v int32) *ListAliasesByKeyIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageSize(v int32) *ListAliasesByKeyIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) Validate() error {
	return dara.Validate(s)
}
