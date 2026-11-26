// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplacePersonalSourceObjectBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetAppToken() *string
	SetObjectBindingsShrink(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetObjectBindingsShrink() *string
	SetOpsUserCode(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest
	GetTenantId() *string
}

type ReplacePersonalSourceObjectBindingsShrinkRequest struct {
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
	// The ID of the data source. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplacePersonalSourceObjectBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetAppToken(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetObjectBindingsShrink(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetOpsUserCode(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetOpsUserId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetSourceId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) SetTenantId(v string) *ReplacePersonalSourceObjectBindingsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
