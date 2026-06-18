// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CheckCloudResourceAuthorizedRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CheckCloudResourceAuthorizedRequest
	GetRegionId() *string
	SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest
	GetRoleArn() *string
}

type CheckCloudResourceAuthorizedRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The global resource descriptor ARN (Alibaba Cloud Resource Name) of the authorized role. After the authorization of this role is complete, the related KMS can be used. Format: acs:ram::$accountID:role/$roleName.
	//
	// - $accountID: the Alibaba Cloud account ID. To view the ID, logon to the Alibaba Cloud Management Console, move the mouse over the profile picture in the upper-right corner, and then click Security Settings.
	//
	// - $roleName: the RAM role name. The value is fixed as AliyunRdsInstanceEncryptionDefaultRole.
	//
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CheckCloudResourceAuthorizedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCloudResourceAuthorizedRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceName(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
