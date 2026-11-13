// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetVersion interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DatasetVersion
	GetComment() *string
	SetCreateTime(v int64) *DatasetVersion
	GetCreateTime() *int64
	SetCreatorId(v string) *DatasetVersion
	GetCreatorId() *string
	SetDatasetId(v string) *DatasetVersion
	GetDatasetId() *string
	SetId(v string) *DatasetVersion
	GetId() *string
	SetImportInfo(v map[string]*string) *DatasetVersion
	GetImportInfo() map[string]*string
	SetLabels(v []*DatasetLabel) *DatasetVersion
	GetLabels() []*DatasetLabel
	SetModifyTime(v int64) *DatasetVersion
	GetModifyTime() *int64
	SetMountPath(v string) *DatasetVersion
	GetMountPath() *string
	SetStorageType(v string) *DatasetVersion
	GetStorageType() *string
	SetUrl(v string) *DatasetVersion
	GetUrl() *string
	SetVersionNumber(v int32) *DatasetVersion
	GetVersionNumber() *int32
}

type DatasetVersion struct {
	// The dataset version description.
	//
	// example:
	//
	// Initial version
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation timestamp in milliseconds.
	//
	// example:
	//
	// 1736756055000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 17815XXX61016173
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The associated dataset ID.
	//
	// example:
	//
	// dataworks-datasetVersion:0gfxxxjx155usz3hrv
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset version ID.
	//
	// example:
	//
	// dataworks-datasetVersion:0gfxxxjx155usz3hrv:1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The storage import configuration of the dataset. The required configuration information varies based on the storage type.
	//
	// <details>
	//
	// <summary>NAS</summary>
	//
	// For valid values, refer to the response of the DescribeFileSystems API operation.
	//
	// ```JSON
	//
	// {
	//
	//   "fileSystemId": "3b6XXX89c9", // File system ID
	//
	//   "fileSystemStorageType": "Performance", // Storage specification of the file system
	//
	//   "vpcId": "vpc-uf66oxxxrqge1t2gson7s" // VPC ID of the mount target
	//
	// }
	//
	// ```
	//
	// </details>
	//
	// example:
	//
	// {"fileSystemId":"3b6XXX89c9","fileSystemStorageType":"Performance","vpcId":"vpc-uf66oxxxrqge1t2gson7s"}
	ImportInfo map[string]*string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// The PAI dataset labels.
	Labels []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The modification timestamp in milliseconds.
	//
	// example:
	//
	// 1736756055000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The mount path. Default value: /mnt/data.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The storage type. This property is read-only and remains consistent with the corresponding property of the associated dataset.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The URL.
	//
	// example:
	//
	// oss://test-oss-bucket/test_dir/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The dataset version number.
	//
	// example:
	//
	// 1
	VersionNumber *int32 `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s DatasetVersion) String() string {
	return dara.Prettify(s)
}

func (s DatasetVersion) GoString() string {
	return s.String()
}

func (s *DatasetVersion) GetComment() *string {
	return s.Comment
}

func (s *DatasetVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DatasetVersion) GetCreatorId() *string {
	return s.CreatorId
}

func (s *DatasetVersion) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DatasetVersion) GetId() *string {
	return s.Id
}

func (s *DatasetVersion) GetImportInfo() map[string]*string {
	return s.ImportInfo
}

func (s *DatasetVersion) GetLabels() []*DatasetLabel {
	return s.Labels
}

func (s *DatasetVersion) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DatasetVersion) GetMountPath() *string {
	return s.MountPath
}

func (s *DatasetVersion) GetStorageType() *string {
	return s.StorageType
}

func (s *DatasetVersion) GetUrl() *string {
	return s.Url
}

func (s *DatasetVersion) GetVersionNumber() *int32 {
	return s.VersionNumber
}

func (s *DatasetVersion) SetComment(v string) *DatasetVersion {
	s.Comment = &v
	return s
}

func (s *DatasetVersion) SetCreateTime(v int64) *DatasetVersion {
	s.CreateTime = &v
	return s
}

func (s *DatasetVersion) SetCreatorId(v string) *DatasetVersion {
	s.CreatorId = &v
	return s
}

func (s *DatasetVersion) SetDatasetId(v string) *DatasetVersion {
	s.DatasetId = &v
	return s
}

func (s *DatasetVersion) SetId(v string) *DatasetVersion {
	s.Id = &v
	return s
}

func (s *DatasetVersion) SetImportInfo(v map[string]*string) *DatasetVersion {
	s.ImportInfo = v
	return s
}

func (s *DatasetVersion) SetLabels(v []*DatasetLabel) *DatasetVersion {
	s.Labels = v
	return s
}

func (s *DatasetVersion) SetModifyTime(v int64) *DatasetVersion {
	s.ModifyTime = &v
	return s
}

func (s *DatasetVersion) SetMountPath(v string) *DatasetVersion {
	s.MountPath = &v
	return s
}

func (s *DatasetVersion) SetStorageType(v string) *DatasetVersion {
	s.StorageType = &v
	return s
}

func (s *DatasetVersion) SetUrl(v string) *DatasetVersion {
	s.Url = &v
	return s
}

func (s *DatasetVersion) SetVersionNumber(v int32) *DatasetVersion {
	s.VersionNumber = &v
	return s
}

func (s *DatasetVersion) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
