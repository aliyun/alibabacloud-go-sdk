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
	// example:
	//
	// cl-0004i0nlngorvgjpl7ck
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// {\\"dataServerAddresses\\":[{\\"host\\":\\"111\\",\\"port\\":\\"\\"}],\\"sharePath\\":\\"/share\\",\\"mountOptions\\":\\"\\",\\"fileSystemType\\":\\"nfs\\"}
	ConnectionInfo *string `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty"`
	// example:
	//
	// {}
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbr_detection_source-nas-20250710101315
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COMMON_NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// example:
	//
	// META
	IndexLevel *string `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	// example:
	//
	// {\\"includes\\":[],\\"excludes\\":[],\\"conflictPolicy\\":\\"OVERWRITE_EXISTING\\"}
	Options *string   `json:"Options,omitempty" xml:"Options,omitempty"`
	Path    []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// example:
	//
	// I|1729493847|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// example:
	//
	// 0:24:5120
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
