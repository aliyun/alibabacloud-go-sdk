// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAvailableConfigsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListAvailableConfigsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAvailableConfigsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListAvailableConfigsRequest
	GetTenantId() *string
}

type ListAvailableConfigsRequest struct {
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
	// The tenant ID. This is a common parameter. For winnexo-cli, it is explicitly passed by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAvailableConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableConfigsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAvailableConfigsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAvailableConfigsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAvailableConfigsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAvailableConfigsRequest) SetAppToken(v string) *ListAvailableConfigsRequest {
	s.AppToken = &v
	return s
}

func (s *ListAvailableConfigsRequest) SetOpsUserCode(v string) *ListAvailableConfigsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAvailableConfigsRequest) SetOpsUserId(v string) *ListAvailableConfigsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAvailableConfigsRequest) SetTenantId(v string) *ListAvailableConfigsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAvailableConfigsRequest) Validate() error {
	return dara.Validate(s)
}
