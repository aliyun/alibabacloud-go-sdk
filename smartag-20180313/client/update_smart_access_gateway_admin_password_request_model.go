// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayAdminPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetCrossAccount() *bool
	SetPassword(v string) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetPassword() *string
	SetRegionId(v string) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetResourceUid() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayAdminPasswordRequest
	GetSagSn() *string
}

type UpdateSmartAccessGatewayAdminPasswordRequest struct {
	// Specifies whether to query only Smart Access Gateway instances that belong to other Alibaba Cloud accounts. Valid values:
	//
	// - **false*	- (default): no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The new logon password for the Smart Access Gateway device.
	//
	// Password format requirements: The password must be 8 to 30 characters in length and can contain uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// > Asterisks (*) are used in the sample value to mask the actual password. This does not mean that asterisks are supported in passwords. Specify the password based on the format requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region where the Smart Access Gateway instance resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the Smart Access Gateway instance belongs.
	//
	// example:
	//
	// 147304382796****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The Smart Access Gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-sv487b7lno6go5****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the Smart Access Gateway device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4ea****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s UpdateSmartAccessGatewayAdminPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayAdminPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetPassword(v string) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetRegionId(v string) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) SetSagSn(v string) *UpdateSmartAccessGatewayAdminPasswordRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordRequest) Validate() error {
	return dara.Validate(s)
}
