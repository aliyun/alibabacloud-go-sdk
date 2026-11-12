// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImageShareAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListCustomImageShareAccountsRequest
	GetClientToken() *string
	SetImageId(v string) *ListCustomImageShareAccountsRequest
	GetImageId() *string
	SetPageNumber(v int32) *ListCustomImageShareAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomImageShareAccountsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomImageShareAccountsRequest
	GetRegionId() *string
}

type ListCustomImageShareAccountsRequest struct {
	// The client token that is used to ensure the idempotence of the request. A client token is generated from your client to ensure uniqueness across different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the shared custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the available regions for simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomImageShareAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImageShareAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCustomImageShareAccountsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListCustomImageShareAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomImageShareAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomImageShareAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomImageShareAccountsRequest) SetClientToken(v string) *ListCustomImageShareAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetImageId(v string) *ListCustomImageShareAccountsRequest {
	s.ImageId = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetPageNumber(v int32) *ListCustomImageShareAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetPageSize(v int32) *ListCustomImageShareAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetRegionId(v string) *ListCustomImageShareAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) Validate() error {
	return dara.Validate(s)
}
