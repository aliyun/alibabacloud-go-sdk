// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountId(v string) *DeleteDatabaseAccountRequest
	GetDatabaseAccountId() *string
	SetInstanceId(v string) *DeleteDatabaseAccountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDatabaseAccountRequest
	GetRegionId() *string
}

type DeleteDatabaseAccountRequest struct {
	// The ID of the database account that you want to delete.
	//
	// >  You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to query the database account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The ID of the bastion host from which you want to delete the database account.
	//
	// > You can call the DescribeInstances operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to delete the database account.
	//
	// > For more information about the mapping between region IDs and region names, [see Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDatabaseAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccountRequest) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *DeleteDatabaseAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDatabaseAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDatabaseAccountRequest) SetDatabaseAccountId(v string) *DeleteDatabaseAccountRequest {
	s.DatabaseAccountId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) SetInstanceId(v string) *DeleteDatabaseAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) SetRegionId(v string) *DeleteDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
