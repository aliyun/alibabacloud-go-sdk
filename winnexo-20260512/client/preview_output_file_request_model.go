// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewOutputFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputItemId(v string) *PreviewOutputFileRequest
	GetOutputItemId() *string
	SetTenantId(v string) *PreviewOutputFileRequest
	GetTenantId() *string
	SetVersionId(v string) *PreviewOutputFileRequest
	GetVersionId() *string
}

type PreviewOutputFileRequest struct {
	// The output item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version referenced by the alias.
	//
	// example:
	//
	// exampleVersionId
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PreviewOutputFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileRequest) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileRequest) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *PreviewOutputFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewOutputFileRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PreviewOutputFileRequest) SetOutputItemId(v string) *PreviewOutputFileRequest {
	s.OutputItemId = &v
	return s
}

func (s *PreviewOutputFileRequest) SetTenantId(v string) *PreviewOutputFileRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewOutputFileRequest) SetVersionId(v string) *PreviewOutputFileRequest {
	s.VersionId = &v
	return s
}

func (s *PreviewOutputFileRequest) Validate() error {
	return dara.Validate(s)
}
