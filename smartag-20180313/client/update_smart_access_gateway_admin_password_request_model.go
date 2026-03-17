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
	// Specifies whether to query only the SAG instances that belong to another Alibaba Cloud account. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The new password used to log on to the SAG device.
	//
	// The password must be 8 to 30 characters in length and can contain letters, digits, and underscores (_).
	//
	// > In the example, asterisks (\\*) are used to conceal the real password. This does not mean that the password supports asterisks (\\*). The actual format requirement prevails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 147304382796****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-sv487b7lno6go5****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
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
