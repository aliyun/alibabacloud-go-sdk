// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *ResetRenderingInstanceRequest
	GetActionName() *string
	SetDataPackageId(v string) *ResetRenderingInstanceRequest
	GetDataPackageId() *string
	SetRenderingInstanceId(v string) *ResetRenderingInstanceRequest
	GetRenderingInstanceId() *string
}

type ResetRenderingInstanceRequest struct {
	// example:
	//
	// Reset
	ActionName    *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ResetRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResetRenderingInstanceRequest) GetActionName() *string {
	return s.ActionName
}

func (s *ResetRenderingInstanceRequest) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *ResetRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ResetRenderingInstanceRequest) SetActionName(v string) *ResetRenderingInstanceRequest {
	s.ActionName = &v
	return s
}

func (s *ResetRenderingInstanceRequest) SetDataPackageId(v string) *ResetRenderingInstanceRequest {
	s.DataPackageId = &v
	return s
}

func (s *ResetRenderingInstanceRequest) SetRenderingInstanceId(v string) *ResetRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ResetRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
