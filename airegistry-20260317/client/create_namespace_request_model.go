// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNamespaceRequest
	GetDescription() *string
	SetName(v string) *CreateNamespaceRequest
	GetName() *string
	SetScanPolicy(v string) *CreateNamespaceRequest
	GetScanPolicy() *string
	SetTags(v string) *CreateNamespaceRequest
	GetTags() *string
}

type CreateNamespaceRequest struct {
	// example:
	//
	// 用于管理客服场景的Prompt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 我的Prompt空间
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"minBlockRiskLevel":"medium","maxSkipRatio":0.2}
	ScanPolicy *string `json:"ScanPolicy,omitempty" xml:"ScanPolicy,omitempty"`
	// example:
	//
	// customer-service,production
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNamespaceRequest) GetScanPolicy() *string {
	return s.ScanPolicy
}

func (s *CreateNamespaceRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateNamespaceRequest) SetDescription(v string) *CreateNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetScanPolicy(v string) *CreateNamespaceRequest {
	s.ScanPolicy = &v
	return s
}

func (s *CreateNamespaceRequest) SetTags(v string) *CreateNamespaceRequest {
	s.Tags = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
