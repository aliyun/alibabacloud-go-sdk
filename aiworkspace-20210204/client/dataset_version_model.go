// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetVersion interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int64) *DatasetVersion
	GetDataCount() *int64
	SetDataSize(v int64) *DatasetVersion
	GetDataSize() *int64
	SetDataSourceType(v string) *DatasetVersion
	GetDataSourceType() *string
	SetDescription(v string) *DatasetVersion
	GetDescription() *string
	SetGmtCreateTime(v string) *DatasetVersion
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *DatasetVersion
	GetGmtModifiedTime() *string
	SetImportInfo(v string) *DatasetVersion
	GetImportInfo() *string
	SetLabels(v []*Label) *DatasetVersion
	GetLabels() []*Label
	SetMountAccess(v string) *DatasetVersion
	GetMountAccess() *string
	SetOptions(v string) *DatasetVersion
	GetOptions() *string
	SetProperty(v string) *DatasetVersion
	GetProperty() *string
	SetSourceId(v string) *DatasetVersion
	GetSourceId() *string
	SetSourceType(v string) *DatasetVersion
	GetSourceType() *string
	SetUri(v string) *DatasetVersion
	GetUri() *string
	SetVersionName(v string) *DatasetVersion
	GetVersionName() *string
}

type DatasetVersion struct {
	// DataCount
	//
	// example:
	//
	// 10000
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// DataSize
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// DataSourceType
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// Description
	//
	// example:
	//
	// base model v1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// create time
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// modify time
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// ImportInfo
	//
	// example:
	//
	// {
	//
	//     "region": "cn-wulanchabu",
	//
	//     "fileSystemId": "bmcpfs-xxxxxxxxxxx",
	//
	//     "path": "/mnt",
	//
	//     "mountTarget": "cpfs-xxxxxxxxxxxx-vpc-gacs9f.cn-wulanchabu.cpfs.aliyuncs.com",
	//
	//     "isVpcMount": true
	//
	// }
	ImportInfo *string `json:"ImportInfo,omitempty" xml:"ImportInfo,omitempty"`
	// Labels
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// MountAccess
	//
	// example:
	//
	// RO
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// Options
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// property
	//
	// example:
	//
	// FILE
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// SourceId
	//
	// example:
	//
	// d-65mrsr5fub4u74lej3
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// SourceType
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Uri
	//
	// example:
	//
	// OSS://xxx
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// version name
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DatasetVersion) String() string {
	return dara.Prettify(s)
}

func (s DatasetVersion) GoString() string {
	return s.String()
}

func (s *DatasetVersion) GetDataCount() *int64 {
	return s.DataCount
}

func (s *DatasetVersion) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DatasetVersion) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DatasetVersion) GetDescription() *string {
	return s.Description
}

func (s *DatasetVersion) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DatasetVersion) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *DatasetVersion) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *DatasetVersion) GetLabels() []*Label {
	return s.Labels
}

func (s *DatasetVersion) GetMountAccess() *string {
	return s.MountAccess
}

func (s *DatasetVersion) GetOptions() *string {
	return s.Options
}

func (s *DatasetVersion) GetProperty() *string {
	return s.Property
}

func (s *DatasetVersion) GetSourceId() *string {
	return s.SourceId
}

func (s *DatasetVersion) GetSourceType() *string {
	return s.SourceType
}

func (s *DatasetVersion) GetUri() *string {
	return s.Uri
}

func (s *DatasetVersion) GetVersionName() *string {
	return s.VersionName
}

func (s *DatasetVersion) SetDataCount(v int64) *DatasetVersion {
	s.DataCount = &v
	return s
}

func (s *DatasetVersion) SetDataSize(v int64) *DatasetVersion {
	s.DataSize = &v
	return s
}

func (s *DatasetVersion) SetDataSourceType(v string) *DatasetVersion {
	s.DataSourceType = &v
	return s
}

func (s *DatasetVersion) SetDescription(v string) *DatasetVersion {
	s.Description = &v
	return s
}

func (s *DatasetVersion) SetGmtCreateTime(v string) *DatasetVersion {
	s.GmtCreateTime = &v
	return s
}

func (s *DatasetVersion) SetGmtModifiedTime(v string) *DatasetVersion {
	s.GmtModifiedTime = &v
	return s
}

func (s *DatasetVersion) SetImportInfo(v string) *DatasetVersion {
	s.ImportInfo = &v
	return s
}

func (s *DatasetVersion) SetLabels(v []*Label) *DatasetVersion {
	s.Labels = v
	return s
}

func (s *DatasetVersion) SetMountAccess(v string) *DatasetVersion {
	s.MountAccess = &v
	return s
}

func (s *DatasetVersion) SetOptions(v string) *DatasetVersion {
	s.Options = &v
	return s
}

func (s *DatasetVersion) SetProperty(v string) *DatasetVersion {
	s.Property = &v
	return s
}

func (s *DatasetVersion) SetSourceId(v string) *DatasetVersion {
	s.SourceId = &v
	return s
}

func (s *DatasetVersion) SetSourceType(v string) *DatasetVersion {
	s.SourceType = &v
	return s
}

func (s *DatasetVersion) SetUri(v string) *DatasetVersion {
	s.Uri = &v
	return s
}

func (s *DatasetVersion) SetVersionName(v string) *DatasetVersion {
	s.VersionName = &v
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
