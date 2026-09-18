// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *PreviewGroupSourceRequest
	GetGroupId() *string
	SetSourceId(v string) *PreviewGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewGroupSourceRequest
	GetTenantId() *string
}

type PreviewGroupSourceRequest struct {
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. You can pass this parameter explicitly by using the --tenant-id option in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PreviewGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewGroupSourceRequest) SetGroupId(v string) *PreviewGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetSourceId(v string) *PreviewGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetTenantId(v string) *PreviewGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
