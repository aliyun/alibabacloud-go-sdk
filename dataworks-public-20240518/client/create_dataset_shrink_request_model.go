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
	// The description of the dataset. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The data type. Valid values:
	//
	// - COMMON: general-purpose (default).
	//
	// - PIC: image.
	//
	// - TEXT: text.
	//
	// - TABLE: table.
	//
	// - VIDEO: video.
	//
	// - AUDIO: audio.
	//
	// - INDEX: index.
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The initial version of the dataset.
	//
	// This parameter is required.
	InitVersionShrink *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	// The name of the dataset. The value must be a non-empty string that is up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origin of the dataset. Only DataWorks is supported.
	//
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The storage type. Valid values:
	//
	// - OSS: Object Storage Service.
	//
	// - NAS: general-purpose NAS file storage.
	//
	// - EXTREMENAS: extreme NAS file storage.
	//
	// - DLF_LANCE: Data Lake Formation.
	//
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
