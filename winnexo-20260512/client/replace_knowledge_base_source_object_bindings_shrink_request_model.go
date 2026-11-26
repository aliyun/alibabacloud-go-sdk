// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetAppToken() *string
	SetObjectBindingsShrink(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetObjectBindingsShrink() *string
	SetOpsUserCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest
	GetTenantId() *string
}

type ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The object bindings.
	//
	// This parameter is required.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The code of the operator.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the data source to be re-parsed. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetAppToken(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetObjectBindingsShrink(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetOpsUserCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetOpsUserId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) SetTenantId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
