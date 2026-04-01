// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportOnlineResizeDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSupportOnlineResizeDiskRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSupportOnlineResizeDiskRequest
	GetOwnerAccount() *string
}

type DescribeSupportOnlineResizeDiskRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeSupportOnlineResizeDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportOnlineResizeDiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportOnlineResizeDiskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSupportOnlineResizeDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSupportOnlineResizeDiskRequest) SetDBInstanceId(v string) *DescribeSupportOnlineResizeDiskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskRequest) SetOwnerAccount(v string) *DescribeSupportOnlineResizeDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskRequest) Validate() error {
	return dara.Validate(s)
}
