// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataset interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Dataset
	GetComment() *string
	SetCreateTime(v int64) *Dataset
	GetCreateTime() *int64
	SetCreatorId(v string) *Dataset
	GetCreatorId() *string
	SetDataType(v string) *Dataset
	GetDataType() *string
	SetId(v string) *Dataset
	GetId() *string
	SetLabels(v []*DatasetLabel) *Dataset
	GetLabels() []*DatasetLabel
	SetLatestVersion(v *DatasetVersion) *Dataset
	GetLatestVersion() *DatasetVersion
	SetModifyTime(v int64) *Dataset
	GetModifyTime() *int64
	SetName(v string) *Dataset
	GetName() *string
	SetOrigin(v string) *Dataset
	GetOrigin() *string
	SetProjectId(v int64) *Dataset
	GetProjectId() *int64
	SetReadme(v string) *Dataset
	GetReadme() *string
	SetStorageType(v string) *Dataset
	GetStorageType() *string
}

type Dataset struct {
	// The description of the dataset. The length cannot exceed 1024 characters.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1736756055000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 210484359
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The data type. Valid values:
	//
	// 	- COMMON
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
	// The dataset ID.
	//
	// example:
	//
	// dataworks-dataset:1gxxxqjx155usz3hrv
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Dataset tags. Supported only for PAI datasets.
	Labels []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest dataset version object.
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The modification time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1736756055000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The dataset name. It must be a non-empty string and cannot exceed 128 characters.
	//
	// example:
	//
	// test_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the dataset. Currently supported sources:
	//
	// 	- DataWorks
	//
	// 	- PAI
	//
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Instructions for using the dataset. Markdown rich text is supported.
	//
	// example:
	//
	// ## introduction
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
	// The storage type. Valid values:
	//
	// 	- OSS: Object Storage Service
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- EXTREM_NAS: Extreme NAS file systems
	//
	// 	- DLF_LANCE: Data Lake Formation
	//
	// 	- CPFS: Cloud Paralleled File System
	//
	// 	- BMCPFS: CPFS for Lingjun
	//
	// 	- MAXCOMPUTE: MaxCompute table
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s Dataset) String() string {
	return dara.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) GetComment() *string {
	return s.Comment
}

func (s *Dataset) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Dataset) GetCreatorId() *string {
	return s.CreatorId
}

func (s *Dataset) GetDataType() *string {
	return s.DataType
}

func (s *Dataset) GetId() *string {
	return s.Id
}

func (s *Dataset) GetLabels() []*DatasetLabel {
	return s.Labels
}

func (s *Dataset) GetLatestVersion() *DatasetVersion {
	return s.LatestVersion
}

func (s *Dataset) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Dataset) GetName() *string {
	return s.Name
}

func (s *Dataset) GetOrigin() *string {
	return s.Origin
}

func (s *Dataset) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *Dataset) GetReadme() *string {
	return s.Readme
}

func (s *Dataset) GetStorageType() *string {
	return s.StorageType
}

func (s *Dataset) SetComment(v string) *Dataset {
	s.Comment = &v
	return s
}

func (s *Dataset) SetCreateTime(v int64) *Dataset {
	s.CreateTime = &v
	return s
}

func (s *Dataset) SetCreatorId(v string) *Dataset {
	s.CreatorId = &v
	return s
}

func (s *Dataset) SetDataType(v string) *Dataset {
	s.DataType = &v
	return s
}

func (s *Dataset) SetId(v string) *Dataset {
	s.Id = &v
	return s
}

func (s *Dataset) SetLabels(v []*DatasetLabel) *Dataset {
	s.Labels = v
	return s
}

func (s *Dataset) SetLatestVersion(v *DatasetVersion) *Dataset {
	s.LatestVersion = v
	return s
}

func (s *Dataset) SetModifyTime(v int64) *Dataset {
	s.ModifyTime = &v
	return s
}

func (s *Dataset) SetName(v string) *Dataset {
	s.Name = &v
	return s
}

func (s *Dataset) SetOrigin(v string) *Dataset {
	s.Origin = &v
	return s
}

func (s *Dataset) SetProjectId(v int64) *Dataset {
	s.ProjectId = &v
	return s
}

func (s *Dataset) SetReadme(v string) *Dataset {
	s.Readme = &v
	return s
}

func (s *Dataset) SetStorageType(v string) *Dataset {
	s.StorageType = &v
	return s
}

func (s *Dataset) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}
