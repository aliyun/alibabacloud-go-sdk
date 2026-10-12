// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int64) *GetDatasetVersionResponseBody
	GetDataCount() *int64
	SetDataSize(v int64) *GetDatasetVersionResponseBody
	GetDataSize() *int64
	SetDataSourceType(v string) *GetDatasetVersionResponseBody
	GetDataSourceType() *string
	SetDatasetId(v string) *GetDatasetVersionResponseBody
	GetDatasetId() *string
	SetDatasetTaskRamRole(v string) *GetDatasetVersionResponseBody
	GetDatasetTaskRamRole() *string
	SetDescription(v string) *GetDatasetVersionResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetDatasetVersionResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetDatasetVersionResponseBody
	GetGmtModifiedTime() *string
	SetImportInfo(v string) *GetDatasetVersionResponseBody
	GetImportInfo() *string
	SetLabels(v []*Label) *GetDatasetVersionResponseBody
	GetLabels() []*Label
	SetMountAccess(v string) *GetDatasetVersionResponseBody
	GetMountAccess() *string
	SetOptions(v string) *GetDatasetVersionResponseBody
	GetOptions() *string
	SetProperty(v string) *GetDatasetVersionResponseBody
	GetProperty() *string
	SetRequestId(v string) *GetDatasetVersionResponseBody
	GetRequestId() *string
	SetSourceId(v string) *GetDatasetVersionResponseBody
	GetSourceId() *string
	SetSourceType(v string) *GetDatasetVersionResponseBody
	GetSourceType() *string
	SetUri(v string) *GetDatasetVersionResponseBody
	GetUri() *string
	SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *GetDatasetVersionResponseBody
	GetUserMetricsEndpoints() []*UserMetricsEndpoint
	SetVersionName(v string) *GetDatasetVersionResponseBody
	GetVersionName() *string
}

type GetDatasetVersionResponseBody struct {
	// The data count.
	//
	// example:
	//
	// 10000
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The dataset size.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data source type. Valid values:
	//
	// - NAS: File Storage NAS.
	//
	// - CPFS: Cloud Parallel File Storage.
	//
	// - OSS: Object Storage Service (OSS).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The primary ID of the resource.
	//
	// example:
	//
	// d-dkdbnnap0g7b6su4yg
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// DatasetTaskRamRole
	//
	// example:
	//
	// acs:ram::1234567890123456:role/role-name
	DatasetTaskRamRole *string `json:"DatasetTaskRamRole,omitempty" xml:"DatasetTaskRamRole,omitempty"`
	// The version description.
	//
	// example:
	//
	// This is a description of the dataset version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ss.sssZ format.
	//
	// example:
	//
	// 2023-12-13T10:22:05.694Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ss.sssZ format.
	//
	// example:
	//
	// 2023-12-13T10:22:05.694Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The storage import configuration of the dataset. OSS, NAS, and CPFS are supported.
	//
	// <details>
	//
	// <summary>OSS</summary>
	//
	// {
	//
	// "region": "${region}",// The region ID
	//
	// "bucket": "${bucket}",// The bucket name
	//
	// "path": "${path}" // The file path
	//
	// }
	//
	// </details>
	//
	// <details>
	//
	// <summary>NAS</summary>
	//
	// </details>
	//
	// <details>
	//
	// <summary>CPFS</summary>
	//
	// Block content
	//
	// </details>
	//
	//
	// <details>
	//
	// <summary>AI CPFS</summary>
	//
	// Block content
	//
	// </details>
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
	// The resource tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The permissions when the dataset is mounted. Valid values:
	//
	// - RO: Read-only mount.
	//
	// - RW: Read/write mount.
	//
	// example:
	//
	// RO
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The extended fields.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The property of the dataset. Valid values:
	//
	// - DIRECTORY: Directory.
	//
	// - FILE: File.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C55DF3DA-F120-5E37-A374-F49365531701
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source ID of the dataset.
	//
	// example:
	//
	// d-rbvg5wzljzjhc9ks92
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. Valid values:
	//
	// - ENHANCE: Data augmentation.
	//
	// - ADDVERSION: Added version.
	//
	// - CLEANING: Data cleaning.
	//
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The URI configuration example.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://ai4d-br7hx9ngzelo2o6uip.oss-cn-shanghai.aliyuncs.com/365349/data-1157703270994901/datasets/aka108o/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The list of user metric collection endpoints.
	//
	// example:
	//
	// [{"Port":0,"Path":""}]
	UserMetricsEndpoints []*UserMetricsEndpoint `json:"UserMetricsEndpoints,omitempty" xml:"UserMetricsEndpoints,omitempty" type:"Repeated"`
	// The dataset version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionResponseBody) GetDataCount() *int64 {
	return s.DataCount
}

func (s *GetDatasetVersionResponseBody) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetDatasetVersionResponseBody) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetDatasetVersionResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetDatasetVersionResponseBody) GetDatasetTaskRamRole() *string {
	return s.DatasetTaskRamRole
}

func (s *GetDatasetVersionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDatasetVersionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetDatasetVersionResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetDatasetVersionResponseBody) GetImportInfo() *string {
	return s.ImportInfo
}

func (s *GetDatasetVersionResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetDatasetVersionResponseBody) GetMountAccess() *string {
	return s.MountAccess
}

func (s *GetDatasetVersionResponseBody) GetOptions() *string {
	return s.Options
}

func (s *GetDatasetVersionResponseBody) GetProperty() *string {
	return s.Property
}

func (s *GetDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetVersionResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetDatasetVersionResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetDatasetVersionResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetDatasetVersionResponseBody) GetUserMetricsEndpoints() []*UserMetricsEndpoint {
	return s.UserMetricsEndpoints
}

func (s *GetDatasetVersionResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetDatasetVersionResponseBody) SetDataCount(v int64) *GetDatasetVersionResponseBody {
	s.DataCount = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDataSize(v int64) *GetDatasetVersionResponseBody {
	s.DataSize = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDataSourceType(v string) *GetDatasetVersionResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDatasetId(v string) *GetDatasetVersionResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDatasetTaskRamRole(v string) *GetDatasetVersionResponseBody {
	s.DatasetTaskRamRole = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDescription(v string) *GetDatasetVersionResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetGmtCreateTime(v string) *GetDatasetVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetGmtModifiedTime(v string) *GetDatasetVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetImportInfo(v string) *GetDatasetVersionResponseBody {
	s.ImportInfo = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetLabels(v []*Label) *GetDatasetVersionResponseBody {
	s.Labels = v
	return s
}

func (s *GetDatasetVersionResponseBody) SetMountAccess(v string) *GetDatasetVersionResponseBody {
	s.MountAccess = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetOptions(v string) *GetDatasetVersionResponseBody {
	s.Options = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetProperty(v string) *GetDatasetVersionResponseBody {
	s.Property = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetRequestId(v string) *GetDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetSourceId(v string) *GetDatasetVersionResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetSourceType(v string) *GetDatasetVersionResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetUri(v string) *GetDatasetVersionResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *GetDatasetVersionResponseBody {
	s.UserMetricsEndpoints = v
	return s
}

func (s *GetDatasetVersionResponseBody) SetVersionName(v string) *GetDatasetVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetDatasetVersionResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserMetricsEndpoints != nil {
		for _, item := range s.UserMetricsEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
