// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeUserEncryptionKeyListRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeUserEncryptionKeyListRequest
	GetRegionId() *string
}

type DescribeUserEncryptionKeyListRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of KMS keys to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeUserEncryptionKeyListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeUserEncryptionKeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) Validate() error {
	return dara.Validate(s)
}
