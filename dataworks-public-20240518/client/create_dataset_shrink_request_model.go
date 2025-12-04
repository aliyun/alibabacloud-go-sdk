// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDatasetShrinkRequest
	GetComment() *string
	SetDataType(v string) *CreateDatasetShrinkRequest
	GetDataType() *string
	SetInitVersionShrink(v string) *CreateDatasetShrinkRequest
	GetInitVersionShrink() *string
	SetName(v string) *CreateDatasetShrinkRequest
	GetName() *string
	SetOrigin(v string) *CreateDatasetShrinkRequest
	GetOrigin() *string
	SetProjectId(v int64) *CreateDatasetShrinkRequest
	GetProjectId() *int64
	SetStorageType(v string) *CreateDatasetShrinkRequest
	GetStorageType() *string
}

type CreateDatasetShrinkRequest struct {
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	InitVersionShrink *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_oss_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatasetShrinkRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDatasetShrinkRequest) GetInitVersionShrink() *string {
	return s.InitVersionShrink
}

func (s *CreateDatasetShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatasetShrinkRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateDatasetShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDatasetShrinkRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDatasetShrinkRequest) SetComment(v string) *CreateDatasetShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDataType(v string) *CreateDatasetShrinkRequest {
	s.DataType = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetInitVersionShrink(v string) *CreateDatasetShrinkRequest {
	s.InitVersionShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetName(v string) *CreateDatasetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetOrigin(v string) *CreateDatasetShrinkRequest {
	s.Origin = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetProjectId(v int64) *CreateDatasetShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetStorageType(v string) *CreateDatasetShrinkRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
