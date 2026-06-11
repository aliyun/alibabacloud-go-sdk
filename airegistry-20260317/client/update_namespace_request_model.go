// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateNamespaceRequest
	GetDescription() *string
	SetName(v string) *UpdateNamespaceRequest
	GetName() *string
	SetNamespaceId(v string) *UpdateNamespaceRequest
	GetNamespaceId() *string
	SetScanPolicy(v string) *UpdateNamespaceRequest
	GetScanPolicy() *string
	SetTags(v string) *UpdateNamespaceRequest
	GetTags() *string
}

type UpdateNamespaceRequest struct {
	// example:
	//
	// 用于管理客服场景的Prompt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 我的Prompt空间
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	ScanPolicy  *string `json:"ScanPolicy,omitempty" xml:"ScanPolicy,omitempty"`
	// example:
	//
	// customer-service,production
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceRequest) GetScanPolicy() *string {
	return s.ScanPolicy
}

func (s *UpdateNamespaceRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateNamespaceRequest) SetDescription(v string) *UpdateNamespaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateNamespaceRequest) SetName(v string) *UpdateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceId(v string) *UpdateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetScanPolicy(v string) *UpdateNamespaceRequest {
	s.ScanPolicy = &v
	return s
}

func (s *UpdateNamespaceRequest) SetTags(v string) *UpdateNamespaceRequest {
	s.Tags = &v
	return s
}

func (s *UpdateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
