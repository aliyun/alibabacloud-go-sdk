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
	// The description of the dataset. It must not exceed 1,024 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The data type. Valid values:
	//
	// 	- COMMON: Common (Default)
	//
	// 	- PIC
	//
	// 	- TEXT
	//
	// 	- TABLE
	//
	// 	- VIDEO
	//
	// 	- AUDIO
	//
	// 	- INDEX
	//
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The initial version of the dataset.
	//
	// This parameter is required.
	InitVersionShrink *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	// The name of the dataset. It cannot be an empty string and must not exceed 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the dataset. Currently, only DataWorks is supported.
	//
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The storage type. Currently supported values:
	//
	// 	- OSS
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- EXTREMENAS: Extreme NAS file systems
	//
	// 	- DLF_LANCE: Data Lake Formation
	//
	// Valid values:
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- MAXCOMPUTE: MaxCompute table
	//
	// 	- CPFS: Cloud Parallel File Storage
	//
	// 	- BMCPFS: CPFS for Lingjun
	//
	// 	- EXTREMENAS: Extreme NAS file systems
	//
	// 	- OSS: Object Storage Service
	//
	// 	- DLF_LANCE: Data Lake Formation.
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
