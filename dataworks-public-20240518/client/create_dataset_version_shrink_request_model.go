// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDatasetVersionShrinkRequest
	GetComment() *string
	SetDatasetId(v string) *CreateDatasetVersionShrinkRequest
	GetDatasetId() *string
	SetImportInfoShrink(v string) *CreateDatasetVersionShrinkRequest
	GetImportInfoShrink() *string
	SetMountPath(v string) *CreateDatasetVersionShrinkRequest
	GetMountPath() *string
	SetUrl(v string) *CreateDatasetVersionShrinkRequest
	GetUrl() *string
}

type CreateDatasetVersionShrinkRequest struct {
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	DatasetId        *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	ImportInfoShrink *string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://test-oss-bucket/test_dir/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDatasetVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetVersionShrinkRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetVersionShrinkRequest) GetImportInfoShrink() *string {
	return s.ImportInfoShrink
}

func (s *CreateDatasetVersionShrinkRequest) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateDatasetVersionShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDatasetVersionShrinkRequest) SetComment(v string) *CreateDatasetVersionShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetDatasetId(v string) *CreateDatasetVersionShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetImportInfoShrink(v string) *CreateDatasetVersionShrinkRequest {
	s.ImportInfoShrink = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetMountPath(v string) *CreateDatasetVersionShrinkRequest {
	s.MountPath = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) SetUrl(v string) *CreateDatasetVersionShrinkRequest {
	s.Url = &v
	return s
}

func (s *CreateDatasetVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
