// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *QueryAppTopologyRequest
	GetAppType() *string
	SetDb(v string) *QueryAppTopologyRequest
	GetDb() *string
	SetDbName(v string) *QueryAppTopologyRequest
	GetDbName() *string
	SetEndTime(v int64) *QueryAppTopologyRequest
	GetEndTime() *int64
	SetFilters(v map[string]*string) *QueryAppTopologyRequest
	GetFilters() map[string]*string
	SetPid(v string) *QueryAppTopologyRequest
	GetPid() *string
	SetRegionId(v string) *QueryAppTopologyRequest
	GetRegionId() *string
	SetRpc(v string) *QueryAppTopologyRequest
	GetRpc() *string
	SetStartTime(v int64) *QueryAppTopologyRequest
	GetStartTime() *int64
	SetType(v string) *QueryAppTopologyRequest
	GetType() *string
}

type QueryAppTopologyRequest struct {
	// The application type
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The database domain name.
	//
	// example:
	//
	// rm-xxx.mysql.rds.aliyuncs.com:3306
	Db *string `json:"Db,omitempty" xml:"Db,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// orders
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671952708499
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	Filters map[string]*string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The ID of the application.
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of an application. The URL in the address bar contains the process ID (PID) of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is eb4zdose6v%409781be0f44d\\*\\*\\*\\*, you must replace %40 with an at sign (@) to obtain eb4zdose6v@9781be0f44d\\*\\*\\*\\*.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// RPC interface name.
	//
	// example:
	//
	// /eventCenter
	Rpc *string `json:"Rpc,omitempty" xml:"Rpc,omitempty"`
	// The start of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595568910000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type kind of topology.
	//
	// This parameter is required.
	//
	// example:
	//
	// apm_apps_v2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryAppTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppTopologyRequest) GoString() string {
	return s.String()
}

func (s *QueryAppTopologyRequest) GetAppType() *string {
	return s.AppType
}

func (s *QueryAppTopologyRequest) GetDb() *string {
	return s.Db
}

func (s *QueryAppTopologyRequest) GetDbName() *string {
	return s.DbName
}

func (s *QueryAppTopologyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryAppTopologyRequest) GetFilters() map[string]*string {
	return s.Filters
}

func (s *QueryAppTopologyRequest) GetPid() *string {
	return s.Pid
}

func (s *QueryAppTopologyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAppTopologyRequest) GetRpc() *string {
	return s.Rpc
}

func (s *QueryAppTopologyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryAppTopologyRequest) GetType() *string {
	return s.Type
}

func (s *QueryAppTopologyRequest) SetAppType(v string) *QueryAppTopologyRequest {
	s.AppType = &v
	return s
}

func (s *QueryAppTopologyRequest) SetDb(v string) *QueryAppTopologyRequest {
	s.Db = &v
	return s
}

func (s *QueryAppTopologyRequest) SetDbName(v string) *QueryAppTopologyRequest {
	s.DbName = &v
	return s
}

func (s *QueryAppTopologyRequest) SetEndTime(v int64) *QueryAppTopologyRequest {
	s.EndTime = &v
	return s
}

func (s *QueryAppTopologyRequest) SetFilters(v map[string]*string) *QueryAppTopologyRequest {
	s.Filters = v
	return s
}

func (s *QueryAppTopologyRequest) SetPid(v string) *QueryAppTopologyRequest {
	s.Pid = &v
	return s
}

func (s *QueryAppTopologyRequest) SetRegionId(v string) *QueryAppTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAppTopologyRequest) SetRpc(v string) *QueryAppTopologyRequest {
	s.Rpc = &v
	return s
}

func (s *QueryAppTopologyRequest) SetStartTime(v int64) *QueryAppTopologyRequest {
	s.StartTime = &v
	return s
}

func (s *QueryAppTopologyRequest) SetType(v string) *QueryAppTopologyRequest {
	s.Type = &v
	return s
}

func (s *QueryAppTopologyRequest) Validate() error {
	return dara.Validate(s)
}
