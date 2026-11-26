// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGraphInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateGraphInfoRequest
	GetAppToken() *string
	SetBusinessProfile(v string) *UpdateGraphInfoRequest
	GetBusinessProfile() *string
	SetDisplayName(v string) *UpdateGraphInfoRequest
	GetDisplayName() *string
	SetGraphName(v string) *UpdateGraphInfoRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *UpdateGraphInfoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateGraphInfoRequest
	GetOpsUserId() *string
	SetTenantId(v string) *UpdateGraphInfoRequest
	GetTenantId() *string
}

type UpdateGraphInfoRequest struct {
	// The application token issued by the console. If you specify this parameter, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The business description of the graph. If this parameter is not specified, the value is an empty string.
	//
	// example:
	//
	// Customer domain semantic graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The display name.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the graph.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The code of the proxy user, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the proxy user, which is the user ID on the WINNEXO platform. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateGraphInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGraphInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateGraphInfoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateGraphInfoRequest) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *UpdateGraphInfoRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateGraphInfoRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *UpdateGraphInfoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateGraphInfoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateGraphInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGraphInfoRequest) SetAppToken(v string) *UpdateGraphInfoRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetBusinessProfile(v string) *UpdateGraphInfoRequest {
	s.BusinessProfile = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetDisplayName(v string) *UpdateGraphInfoRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetGraphName(v string) *UpdateGraphInfoRequest {
	s.GraphName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetOpsUserCode(v string) *UpdateGraphInfoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetOpsUserId(v string) *UpdateGraphInfoRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetTenantId(v string) *UpdateGraphInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGraphInfoRequest) Validate() error {
	return dara.Validate(s)
}
