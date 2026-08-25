// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddDataSourceRequest
	GetClusterId() *string
	SetConnectionInfo(v string) *AddDataSourceRequest
	GetConnectionInfo() *string
	SetCredential(v string) *AddDataSourceRequest
	GetCredential() *string
	SetDataSourceName(v string) *AddDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *AddDataSourceRequest
	GetDataSourceType() *string
	SetExclude(v string) *AddDataSourceRequest
	GetExclude() *string
	SetInclude(v string) *AddDataSourceRequest
	GetInclude() *string
	SetIndexLevel(v string) *AddDataSourceRequest
	GetIndexLevel() *string
	SetOptions(v string) *AddDataSourceRequest
	GetOptions() *string
	SetPath(v []*string) *AddDataSourceRequest
	GetPath() []*string
	SetSchedule(v string) *AddDataSourceRequest
	GetSchedule() *string
	SetSpeedLimit(v string) *AddDataSourceRequest
	GetSpeedLimit() *string
}

type AddDataSourceRequest struct {
	// The ID of the client group used to access the data source.
	//
	// example:
	//
	// cl-0004i0nlngorvgjpl7ck
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection information for the data source. Examples:
	//
	// - Local NAS data source:
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
	// - CPFS AI-Computing Edition data source:
	//
	//   {"vpcMountTarget":"cpfs-010wn\\*\\*\\*wy-vpc-ta\\*\\*\\*8.cn-shanghai.cpfs.aliyuncs.com","sharePath":"/"}
	//
	// - Other large-scale file system data source:
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
	// The access credential for the data source. This parameter is required for local NAS (SMB) data sources and for OSS- or S3-compatible data sources. Examples:
	//
	// - Local NAS data source (SMB protocol):
	//
	//   {
	//
	//   "mountUsername": "\\*\\*\\*\\*\\*",
	//
	//   "mountPassword": "\\*\\*\\*\\*\\*"
	//
	//   }
	//
	// - OSS-compatible or S3-compatible data source:
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
	// {
	//
	//     "mountUsername": "*****",
	//
	//     "mountPassword": "*****"
	//
	// }
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyLocalNas
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The type of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// COMMON_NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// A filter that specifies the files to exclude. This parameter applies only to data source analysis with the archive feature.
	//
	// example:
	//
	// ["/home/alice/log"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// A filter that specifies the files to include. This parameter applies only to data source analysis with the archive feature.
	//
	// example:
	//
	// ["/home/alice", "/home/bob"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The index level. This parameter applies only to data source analysis with the archive feature.
	//
	// - `OFF`: No index is created.
	//
	// - `META`: A metadata index is created.
	//
	// - `ALL`: A full-text index is created. (Deprecated)
	//
	// example:
	//
	// META
	IndexLevel *string `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	// The options for data source analysis. This parameter applies only to data source analysis with the archive feature.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// A list of paths for data source analysis. This parameter applies only to data source analysis with the archive feature.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// The execution schedule. This parameter applies only to data source analysis with the archive feature. The format `I|{startTime}|{interval}` specifies a task that runs at a defined interval starting from `{startTime}`. `{interval}` is an ISO 8601 duration. For example, `PT1H` specifies a one-hour interval, and `P1D` specifies a one-day interval.
	//
	// example:
	//
	// I|1729493847|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The rate limiting settings. This parameter applies only to data source analysis with the archive feature.
	//
	// example:
	//
	// 6:21:10240
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
}

func (s AddDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddDataSourceRequest) GetConnectionInfo() *string {
	return s.ConnectionInfo
}

func (s *AddDataSourceRequest) GetCredential() *string {
	return s.Credential
}

func (s *AddDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *AddDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *AddDataSourceRequest) GetExclude() *string {
	return s.Exclude
}

func (s *AddDataSourceRequest) GetInclude() *string {
	return s.Include
}

func (s *AddDataSourceRequest) GetIndexLevel() *string {
	return s.IndexLevel
}

func (s *AddDataSourceRequest) GetOptions() *string {
	return s.Options
}

func (s *AddDataSourceRequest) GetPath() []*string {
	return s.Path
}

func (s *AddDataSourceRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *AddDataSourceRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *AddDataSourceRequest) SetClusterId(v string) *AddDataSourceRequest {
	s.ClusterId = &v
	return s
}

func (s *AddDataSourceRequest) SetConnectionInfo(v string) *AddDataSourceRequest {
	s.ConnectionInfo = &v
	return s
}

func (s *AddDataSourceRequest) SetCredential(v string) *AddDataSourceRequest {
	s.Credential = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceName(v string) *AddDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceType(v string) *AddDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *AddDataSourceRequest) SetExclude(v string) *AddDataSourceRequest {
	s.Exclude = &v
	return s
}

func (s *AddDataSourceRequest) SetInclude(v string) *AddDataSourceRequest {
	s.Include = &v
	return s
}

func (s *AddDataSourceRequest) SetIndexLevel(v string) *AddDataSourceRequest {
	s.IndexLevel = &v
	return s
}

func (s *AddDataSourceRequest) SetOptions(v string) *AddDataSourceRequest {
	s.Options = &v
	return s
}

func (s *AddDataSourceRequest) SetPath(v []*string) *AddDataSourceRequest {
	s.Path = v
	return s
}

func (s *AddDataSourceRequest) SetSchedule(v string) *AddDataSourceRequest {
	s.Schedule = &v
	return s
}

func (s *AddDataSourceRequest) SetSpeedLimit(v string) *AddDataSourceRequest {
	s.SpeedLimit = &v
	return s
}

func (s *AddDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
