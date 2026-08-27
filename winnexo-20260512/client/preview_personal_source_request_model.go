// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *PreviewPersonalSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewPersonalSourceRequest
	GetTenantId() *string
}

type PreviewPersonalSourceRequest struct {
	// The data source ID, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. The winnexo-cli passes this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewPersonalSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewPersonalSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewPersonalSourceRequest) SetSourceId(v string) *PreviewPersonalSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetTenantId(v string) *PreviewPersonalSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) Validate() error {
	return dara.Validate(s)
}
