// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *ModifyRenderingInstanceRequest
	GetRenderingInstanceId() *string
	SetRenderingSpec(v string) *ModifyRenderingInstanceRequest
	GetRenderingSpec() *string
	SetStorageSize(v string) *ModifyRenderingInstanceRequest
	GetStorageSize() *string
}

type ModifyRenderingInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	// example:
	//
	// 20
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s ModifyRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingInstanceRequest) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *ModifyRenderingInstanceRequest) GetStorageSize() *string {
	return s.StorageSize
}

func (s *ModifyRenderingInstanceRequest) SetRenderingInstanceId(v string) *ModifyRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingInstanceRequest) SetRenderingSpec(v string) *ModifyRenderingInstanceRequest {
	s.RenderingSpec = &v
	return s
}

func (s *ModifyRenderingInstanceRequest) SetStorageSize(v string) *ModifyRenderingInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *ModifyRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
