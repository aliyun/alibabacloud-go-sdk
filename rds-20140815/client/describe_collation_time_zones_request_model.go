// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollationTimeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeCollationTimeZonesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeCollationTimeZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCollationTimeZonesRequest
	GetResourceOwnerId() *int64
}

type DescribeCollationTimeZonesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCollationTimeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollationTimeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCollationTimeZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCollationTimeZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCollationTimeZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCollationTimeZonesRequest) SetOwnerId(v int64) *DescribeCollationTimeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCollationTimeZonesRequest) SetResourceOwnerAccount(v string) *DescribeCollationTimeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCollationTimeZonesRequest) SetResourceOwnerId(v int64) *DescribeCollationTimeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCollationTimeZonesRequest) Validate() error {
	return dara.Validate(s)
}
