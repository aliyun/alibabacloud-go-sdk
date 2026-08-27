// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateSourceContentRequest
	GetForceSync() *bool
	SetSourceId(v string) *UpdateSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateSourceContentRequest
	GetTenantId() *string
}

type UpdateSourceContentRequest struct {
	// The returned content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether to force synchronous processing.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateSourceContentRequest) SetContent(v string) *UpdateSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateSourceContentRequest) SetForceSync(v bool) *UpdateSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateSourceContentRequest) SetSourceId(v string) *UpdateSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateSourceContentRequest) SetTenantId(v string) *UpdateSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
