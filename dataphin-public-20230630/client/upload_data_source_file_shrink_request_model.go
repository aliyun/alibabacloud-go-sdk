// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSourceFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UploadDataSourceFileShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UploadDataSourceFileShrinkRequest
	GetOpUserId() *string
	SetUploadCommandShrink(v string) *UploadDataSourceFileShrinkRequest
	GetUploadCommandShrink() *string
}

type UploadDataSourceFileShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The request object for uploading a datasource authentication file.
	//
	// This parameter is required.
	UploadCommandShrink *string `json:"UploadCommand,omitempty" xml:"UploadCommand,omitempty"`
}

func (s UploadDataSourceFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSourceFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSourceFileShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UploadDataSourceFileShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UploadDataSourceFileShrinkRequest) GetUploadCommandShrink() *string {
	return s.UploadCommandShrink
}

func (s *UploadDataSourceFileShrinkRequest) SetOpTenantId(v int64) *UploadDataSourceFileShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UploadDataSourceFileShrinkRequest) SetOpUserId(v string) *UploadDataSourceFileShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UploadDataSourceFileShrinkRequest) SetUploadCommandShrink(v string) *UploadDataSourceFileShrinkRequest {
	s.UploadCommandShrink = &v
	return s
}

func (s *UploadDataSourceFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
