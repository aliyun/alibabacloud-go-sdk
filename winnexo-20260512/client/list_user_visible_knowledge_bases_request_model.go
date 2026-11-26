// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListUserVisibleKnowledgeBasesRequest
	GetAppToken() *string
	SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListUserVisibleKnowledgeBasesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListUserVisibleKnowledgeBasesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest
	GetTenantId() *string
}

type ListUserVisibleKnowledgeBasesRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed with the identity of the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The keyword for fuzzy search of component data in the form.
	//
	// example:
	//
	// Product knowledge
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
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
	// The tenant ID. This is a common parameter and is explicitly passed by winnexo-cli through --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetAppToken(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.AppToken = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetOpsUserCode(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetOpsUserId(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
