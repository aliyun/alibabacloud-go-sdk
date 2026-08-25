// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateDataSourceRequest
	GetClusterId() *string
	SetConnectionInfo(v string) *UpdateDataSourceRequest
	GetConnectionInfo() *string
	SetCredential(v string) *UpdateDataSourceRequest
	GetCredential() *string
	SetDataSourceId(v string) *UpdateDataSourceRequest
	GetDataSourceId() *string
	SetDataSourceName(v string) *UpdateDataSourceRequest
	GetDataSourceName() *string
	SetExclude(v string) *UpdateDataSourceRequest
	GetExclude() *string
	SetInclude(v string) *UpdateDataSourceRequest
	GetInclude() *string
	SetIndexLevel(v string) *UpdateDataSourceRequest
	GetIndexLevel() *string
	SetOptions(v string) *UpdateDataSourceRequest
	GetOptions() *string
	SetPath(v []*string) *UpdateDataSourceRequest
	GetPath() []*string
	SetSchedule(v string) *UpdateDataSourceRequest
	GetSchedule() *string
	SetSpeedLimit(v string) *UpdateDataSourceRequest
	GetSpeedLimit() *string
}

type UpdateDataSourceRequest struct {
	// The ID of the client group used to access the data source.
	//
	// example:
	//
	// cl-000**************hg9
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection information for the data source. Examples:
	//
	// - On-premises NAS data source:
	//
	//   {
	//
	//   "dataServerAddresses": [
	//
	//   {
	//
	//   "host": "123.123.123.123",
	//
	//   "port": "8080"
	//
	//   }
	//
	//   ],
	//
	//   "sharePath": "/share",
	//
	//   "mountOptions": "vers=3",
	//
	//   "fileSystemType": "nfs"
	//
	//   }
	//
	// - Intelligent Computing CPFS data source:
	//
	//   {"vpcMountTarget":"cpfs-010wn\\*\\*\\*wy-vpc-ta\\*\\*\\*8.cn-shanghai.cpfs.aliyuncs.com","sharePath":"/"}
	//
	// - Other large-scale file system data sources:
	//
	//   {"path":"/mnt"}
	//
	// example:
	//
	// {
	//
	//     "dataServerAddresses": [
	//
	//         {
	//
	//             "host": "123.123.123.123",
	//
	//             "port": "8080"
	//
	//         }
	//
	//     ],
	//
	//     "sharePath": "/share",
	//
	//     "mountOptions": "vers=3",
	//
	//     "fileSystemType": "nfs"
	//
	// }
	ConnectionInfo *string `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty"`
	// The access credentials for the data source. This parameter is used for on-premises NAS data sources that use the SMB protocol, and for OSS and S3 protocol-compatible data sources. Examples:
	//
	// - On-premises NAS data source (SMB protocol):
	//
	//   {
	//
	//   "mountUsername": "\\*\\*\\*\\*\\*",
	//
	//   "mountPassword": "\\*\\*\\*\\*\\*"
	//
	//   }
	//
	// - OSS protocol-compatible data source/S3 protocol-compatible data source:
	//
	//   {
	//
	//   "accessKeyId": "\\*\\*\\*\\*\\*",
	//
	//   "accessKeySecret": "\\*\\*\\*\\*\\*"
	//
	//   }
	//
	// example:
	//
	// {"mountUsername":"***","mountPassword":"***"}
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-0006xo****dtle
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// MyLocalNas
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// A filter to specify which files to exclude. This parameter applies only to the archive feature.
	//
	// example:
	//
	// ["/home/alice/log/"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// A filter to specify which files to include. This parameter applies only to the archive feature.
	//
	// example:
	//
	// ["/home/alice/", "/home/bob/"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The index level for data source analysis. This parameter applies only to the archive feature.
	//
	// - OFF: No index is created.
	//
	// - META: A metadata index is created.
	//
	// - ALL: A full-text index is created. (Deprecated)
	//
	// example:
	//
	// META
	IndexLevel *string `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	// The options for data source analysis. This parameter applies only to the archive feature.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// A list of paths for data source analysis. This parameter applies only to the archive feature.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// The schedule for data source analysis. This parameter applies only to the archive feature. The format is `I|{startTime}|{interval}`. This specifies a task that starts at `{startTime}` and repeats at the specified `{interval}`. `startTime` is a Unix time value in seconds. `interval` is an ISO 8601 time interval. For example, `PT1H` indicates a one-hour interval, and `P1D` indicates a one-day interval.
	//
	// example:
	//
	// I|1729493847|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The rate limiting configuration for data source analysis. This parameter applies only to the archive feature.
	//
	// example:
	//
	// 0:24:102400
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateDataSourceRequest) GetConnectionInfo() *string {
	return s.ConnectionInfo
}

func (s *UpdateDataSourceRequest) GetCredential() *string {
	return s.Credential
}

func (s *UpdateDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDataSourceRequest) GetExclude() *string {
	return s.Exclude
}

func (s *UpdateDataSourceRequest) GetInclude() *string {
	return s.Include
}

func (s *UpdateDataSourceRequest) GetIndexLevel() *string {
	return s.IndexLevel
}

func (s *UpdateDataSourceRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateDataSourceRequest) GetPath() []*string {
	return s.Path
}

func (s *UpdateDataSourceRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateDataSourceRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *UpdateDataSourceRequest) SetClusterId(v string) *UpdateDataSourceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetConnectionInfo(v string) *UpdateDataSourceRequest {
	s.ConnectionInfo = &v
	return s
}

func (s *UpdateDataSourceRequest) SetCredential(v string) *UpdateDataSourceRequest {
	s.Credential = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceId(v string) *UpdateDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceName(v string) *UpdateDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDataSourceRequest) SetExclude(v string) *UpdateDataSourceRequest {
	s.Exclude = &v
	return s
}

func (s *UpdateDataSourceRequest) SetInclude(v string) *UpdateDataSourceRequest {
	s.Include = &v
	return s
}

func (s *UpdateDataSourceRequest) SetIndexLevel(v string) *UpdateDataSourceRequest {
	s.IndexLevel = &v
	return s
}

func (s *UpdateDataSourceRequest) SetOptions(v string) *UpdateDataSourceRequest {
	s.Options = &v
	return s
}

func (s *UpdateDataSourceRequest) SetPath(v []*string) *UpdateDataSourceRequest {
	s.Path = v
	return s
}

func (s *UpdateDataSourceRequest) SetSchedule(v string) *UpdateDataSourceRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateDataSourceRequest) SetSpeedLimit(v string) *UpdateDataSourceRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
