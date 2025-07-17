// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeNotificationTypesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeNotificationTypesRequest
	GetResourceOwnerAccount() *string
}

type DescribeNotificationTypesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeNotificationTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNotificationTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNotificationTypesRequest) SetOwnerId(v int64) *DescribeNotificationTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNotificationTypesRequest) SetResourceOwnerAccount(v string) *DescribeNotificationTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNotificationTypesRequest) Validate() error {
	return dara.Validate(s)
}
