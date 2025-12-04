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
	Comment       *string         `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime    *int64          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId     *string         `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DataType      *string         `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Id            *string         `json:"Id,omitempty" xml:"Id,omitempty"`
	Labels        []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModifyTime    *int64          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name          *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	Origin        *string         `json:"Origin,omitempty" xml:"Origin,omitempty"`
	ProjectId     *int64          `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Readme        *string         `json:"Readme,omitempty" xml:"Readme,omitempty"`
	StorageType   *string         `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
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
