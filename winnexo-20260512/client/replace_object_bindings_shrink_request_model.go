// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceObjectBindingsShrinkRequest
	GetAppToken() *string
	SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest
	GetObjectBindingsShrink() *string
	SetOpsUserCode(v string) *ReplaceObjectBindingsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceObjectBindingsShrinkRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest
	GetTenantId() *string
}

type ReplaceObjectBindingsShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new list of object bindings for full replacement. Passing an empty list clears all bindings.
	//
	// This parameter is required.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
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
	// The ID of the personal FILE data source to be replaced, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter and is explicitly passed by winnexo-cli through the --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceObjectBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceObjectBindingsShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *ReplaceObjectBindingsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceObjectBindingsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceObjectBindingsShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceObjectBindingsShrinkRequest) SetAppToken(v string) *ReplaceObjectBindingsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetOpsUserCode(v string) *ReplaceObjectBindingsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetOpsUserId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
