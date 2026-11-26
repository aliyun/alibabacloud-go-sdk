// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListGraphsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListGraphsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListGraphsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListGraphsRequest
	GetTenantId() *string
}

type ListGraphsRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphsRequest) GoString() string {
	return s.String()
}

func (s *ListGraphsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListGraphsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListGraphsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListGraphsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphsRequest) SetAppToken(v string) *ListGraphsRequest {
	s.AppToken = &v
	return s
}

func (s *ListGraphsRequest) SetOpsUserCode(v string) *ListGraphsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListGraphsRequest) SetOpsUserId(v string) *ListGraphsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListGraphsRequest) SetTenantId(v string) *ListGraphsRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphsRequest) Validate() error {
	return dara.Validate(s)
}
