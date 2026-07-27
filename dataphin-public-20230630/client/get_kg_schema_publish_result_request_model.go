// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgSchemaPublishResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetKgSchemaPublishResultRequest
	GetOpTenantId() *int64
	SetVersionId(v int32) *GetKgSchemaPublishResultRequest
	GetVersionId() *int32
	SetWorkspaceId(v string) *GetKgSchemaPublishResultRequest
	GetWorkspaceId() *string
}

type GetKgSchemaPublishResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 1
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKgSchemaPublishResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgSchemaPublishResultRequest) GoString() string {
	return s.String()
}

func (s *GetKgSchemaPublishResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgSchemaPublishResultRequest) GetVersionId() *int32 {
	return s.VersionId
}

func (s *GetKgSchemaPublishResultRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgSchemaPublishResultRequest) SetOpTenantId(v int64) *GetKgSchemaPublishResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgSchemaPublishResultRequest) SetVersionId(v int32) *GetKgSchemaPublishResultRequest {
	s.VersionId = &v
	return s
}

func (s *GetKgSchemaPublishResultRequest) SetWorkspaceId(v string) *GetKgSchemaPublishResultRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgSchemaPublishResultRequest) Validate() error {
	return dara.Validate(s)
}
