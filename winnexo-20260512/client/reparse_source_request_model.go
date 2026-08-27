// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceSync(v bool) *ReparseSourceRequest
	GetForceSync() *bool
	SetSourceId(v string) *ReparseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseSourceRequest
	GetTenantId() *string
}

type ReparseSourceRequest struct {
	// Specifies whether to synchronously wait for the re-parsing to complete. Default value: false, which indicates asynchronous queuing.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The ID of the data source to re-parse. This ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseSourceRequest) SetForceSync(v bool) *ReparseSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseSourceRequest) SetSourceId(v string) *ReparseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseSourceRequest) SetTenantId(v string) *ReparseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseSourceRequest) Validate() error {
	return dara.Validate(s)
}
