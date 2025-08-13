// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRuleReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GenerateConfigRuleReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GenerateConfigRuleReportRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GenerateConfigRuleReportRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GenerateConfigRuleReportRequest
	GetResourceOwnerAccount() *string
	SetTargetId(v string) *GenerateConfigRuleReportRequest
	GetTargetId() *string
	SetTargetType(v string) *GenerateConfigRuleReportRequest
	GetTargetType() *string
	SetUserType(v string) *GenerateConfigRuleReportRequest
	GetUserType() *string
}

type GenerateConfigRuleReportRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// 154950938137****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// 	- USER: single-account mode
	//
	// 	- RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// >  This parameter is required if the management account of your resource directory is used to enable the Tag Policy feature in both single-account mode and multi-account mode. The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// USER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GenerateConfigRuleReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRuleReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GenerateConfigRuleReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GenerateConfigRuleReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateConfigRuleReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GenerateConfigRuleReportRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GenerateConfigRuleReportRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GenerateConfigRuleReportRequest) GetUserType() *string {
	return s.UserType
}

func (s *GenerateConfigRuleReportRequest) SetOwnerAccount(v string) *GenerateConfigRuleReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetOwnerId(v int64) *GenerateConfigRuleReportRequest {
	s.OwnerId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetRegionId(v string) *GenerateConfigRuleReportRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetResourceOwnerAccount(v string) *GenerateConfigRuleReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetTargetId(v string) *GenerateConfigRuleReportRequest {
	s.TargetId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetTargetType(v string) *GenerateConfigRuleReportRequest {
	s.TargetType = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetUserType(v string) *GenerateConfigRuleReportRequest {
	s.UserType = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) Validate() error {
	return dara.Validate(s)
}
