// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenType(v string) *GetPolicyEnableStatusRequest
	GetOpenType() *string
	SetOwnerAccount(v string) *GetPolicyEnableStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetPolicyEnableStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetPolicyEnableStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetPolicyEnableStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetPolicyEnableStatusRequest
	GetResourceOwnerId() *string
	SetUserType(v string) *GetPolicyEnableStatusRequest
	GetUserType() *string
}

type GetPolicyEnableStatusRequest struct {
	// The enabling type. Valid values:
	//
	// 	- TAG_POLICY: the Tag Policy feature.
	//
	// 	- VERIFY_NO_TAG: the strong verification feature.
	//
	// example:
	//
	// TAG_POLICY
	OpenType     *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mode of the Tag Policy feature. This parameter specifies a filter condition for the query. Valid values:
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
	// RD
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusRequest) GetOpenType() *string {
	return s.OpenType
}

func (s *GetPolicyEnableStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetPolicyEnableStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPolicyEnableStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyEnableStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPolicyEnableStatusRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetPolicyEnableStatusRequest) GetUserType() *string {
	return s.UserType
}

func (s *GetPolicyEnableStatusRequest) SetOpenType(v string) *GetPolicyEnableStatusRequest {
	s.OpenType = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetOwnerAccount(v string) *GetPolicyEnableStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetOwnerId(v int64) *GetPolicyEnableStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetRegionId(v string) *GetPolicyEnableStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetResourceOwnerAccount(v string) *GetPolicyEnableStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetResourceOwnerId(v string) *GetPolicyEnableStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetUserType(v string) *GetPolicyEnableStatusRequest {
	s.UserType = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
