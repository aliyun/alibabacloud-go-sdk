// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTmpIdentityForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPurpose(v string) *GetUserTmpIdentityForPartnerRequest
	GetAuthPurpose() *string
	SetBizId(v string) *GetUserTmpIdentityForPartnerRequest
	GetBizId() *string
	SetExtend(v string) *GetUserTmpIdentityForPartnerRequest
	GetExtend() *string
	SetServiceLinkedRole(v string) *GetUserTmpIdentityForPartnerRequest
	GetServiceLinkedRole() *string
	SetUserId(v string) *GetUserTmpIdentityForPartnerRequest
	GetUserId() *string
}

type GetUserTmpIdentityForPartnerRequest struct {
	// The purpose of the authorization.
	//
	// example:
	//
	// BindDomain
	AuthPurpose *string `json:"AuthPurpose,omitempty" xml:"AuthPurpose,omitempty"`
	// The business ID of the customer.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Specifies whether fuzzy match is supported for port numbers. Set this parameter to **1*	- to enable fuzzy match. Other values or an empty value indicate that fuzzy match is not supported.
	//
	// example:
	//
	// {\\"deliveryNodeName\\":\\"视觉设计确认\\",\\"deliveryNodeStatus\\":\\"Reject\\",\\"deliveryOperatorRole\\":\\"Customer\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The service-linked role.
	//
	// > For information about the service-linked roles supported by ApsaraDB RDS, see [Service-linked roles](https://help.aliyun.com/document_detail/342840.html).
	//
	// example:
	//
	// AliyunServiceRoleForSasCspm
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1231331311
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserTmpIdentityForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserTmpIdentityForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetUserTmpIdentityForPartnerRequest) GetAuthPurpose() *string {
	return s.AuthPurpose
}

func (s *GetUserTmpIdentityForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetUserTmpIdentityForPartnerRequest) GetExtend() *string {
	return s.Extend
}

func (s *GetUserTmpIdentityForPartnerRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *GetUserTmpIdentityForPartnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserTmpIdentityForPartnerRequest) SetAuthPurpose(v string) *GetUserTmpIdentityForPartnerRequest {
	s.AuthPurpose = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerRequest) SetBizId(v string) *GetUserTmpIdentityForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerRequest) SetExtend(v string) *GetUserTmpIdentityForPartnerRequest {
	s.Extend = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerRequest) SetServiceLinkedRole(v string) *GetUserTmpIdentityForPartnerRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerRequest) SetUserId(v string) *GetUserTmpIdentityForPartnerRequest {
	s.UserId = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
