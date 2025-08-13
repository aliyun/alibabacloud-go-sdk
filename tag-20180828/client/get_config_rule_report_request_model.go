// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetConfigRuleReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetConfigRuleReportRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetConfigRuleReportRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetConfigRuleReportRequest
	GetResourceOwnerAccount() *string
	SetTargetId(v string) *GetConfigRuleReportRequest
	GetTargetId() *string
	SetTargetType(v string) *GetConfigRuleReportRequest
	GetTargetType() *string
	SetUserType(v string) *GetConfigRuleReportRequest
	GetUserType() *string
}

type GetConfigRuleReportRequest struct {
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
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// 	- USER: single-account mode
	//
	// 	- RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// USER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetConfigRuleReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleReportRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetConfigRuleReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetConfigRuleReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConfigRuleReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetConfigRuleReportRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GetConfigRuleReportRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GetConfigRuleReportRequest) GetUserType() *string {
	return s.UserType
}

func (s *GetConfigRuleReportRequest) SetOwnerAccount(v string) *GetConfigRuleReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetOwnerId(v int64) *GetConfigRuleReportRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetRegionId(v string) *GetConfigRuleReportRequest {
	s.RegionId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetResourceOwnerAccount(v string) *GetConfigRuleReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetTargetId(v string) *GetConfigRuleReportRequest {
	s.TargetId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetTargetType(v string) *GetConfigRuleReportRequest {
	s.TargetType = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetUserType(v string) *GetConfigRuleReportRequest {
	s.UserType = &v
	return s
}

func (s *GetConfigRuleReportRequest) Validate() error {
	return dara.Validate(s)
}
