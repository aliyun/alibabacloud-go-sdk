// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverRenderingDataPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPackageId(v string) *RecoverRenderingDataPackageRequest
	GetDataPackageId() *string
	SetLoadMode(v string) *RecoverRenderingDataPackageRequest
	GetLoadMode() *string
	SetRenderingInstanceId(v string) *RecoverRenderingDataPackageRequest
	GetRenderingInstanceId() *string
}

type RecoverRenderingDataPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dp-449ea3d16c0841b8bf33ec5bbc86a152
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	LoadMode      *string `json:"LoadMode,omitempty" xml:"LoadMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RecoverRenderingDataPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverRenderingDataPackageRequest) GoString() string {
	return s.String()
}

func (s *RecoverRenderingDataPackageRequest) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *RecoverRenderingDataPackageRequest) GetLoadMode() *string {
	return s.LoadMode
}

func (s *RecoverRenderingDataPackageRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RecoverRenderingDataPackageRequest) SetDataPackageId(v string) *RecoverRenderingDataPackageRequest {
	s.DataPackageId = &v
	return s
}

func (s *RecoverRenderingDataPackageRequest) SetLoadMode(v string) *RecoverRenderingDataPackageRequest {
	s.LoadMode = &v
	return s
}

func (s *RecoverRenderingDataPackageRequest) SetRenderingInstanceId(v string) *RecoverRenderingDataPackageRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *RecoverRenderingDataPackageRequest) Validate() error {
	return dara.Validate(s)
}
