// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceSync(v bool) *ReparseGroupSourceRequest
	GetForceSync() *bool
	SetGroupId(v string) *ReparseGroupSourceRequest
	GetGroupId() *string
	SetSourceId(v string) *ReparseGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseGroupSourceRequest
	GetTenantId() *string
}

type ReparseGroupSourceRequest struct {
	// Specifies whether to synchronously wait for the re-parsing to complete. Default value: false, which indicates that the request is asynchronously queued.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseGroupSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReparseGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseGroupSourceRequest) SetForceSync(v bool) *ReparseGroupSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetGroupId(v string) *ReparseGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetSourceId(v string) *ReparseGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetTenantId(v string) *ReparseGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
