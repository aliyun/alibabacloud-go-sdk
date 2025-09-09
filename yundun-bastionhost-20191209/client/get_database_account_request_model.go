// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountId(v string) *GetDatabaseAccountRequest
	GetDatabaseAccountId() *string
	SetInstanceId(v string) *GetDatabaseAccountRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetDatabaseAccountRequest
	GetRegionId() *string
}

type GetDatabaseAccountRequest struct {
	// The ID of the database account to query.
	//
	// >  You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to query the database account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-i7m2d7zrw11
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDatabaseAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseAccountRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseAccountRequest) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *GetDatabaseAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDatabaseAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDatabaseAccountRequest) SetDatabaseAccountId(v string) *GetDatabaseAccountRequest {
	s.DatabaseAccountId = &v
	return s
}

func (s *GetDatabaseAccountRequest) SetInstanceId(v string) *GetDatabaseAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDatabaseAccountRequest) SetRegionId(v string) *GetDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
