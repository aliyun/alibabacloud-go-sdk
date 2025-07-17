// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppTopologyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *QueryAppTopologyShrinkRequest
	GetAppType() *string
	SetDb(v string) *QueryAppTopologyShrinkRequest
	GetDb() *string
	SetDbName(v string) *QueryAppTopologyShrinkRequest
	GetDbName() *string
	SetEndTime(v int64) *QueryAppTopologyShrinkRequest
	GetEndTime() *int64
	SetFiltersShrink(v string) *QueryAppTopologyShrinkRequest
	GetFiltersShrink() *string
	SetPid(v string) *QueryAppTopologyShrinkRequest
	GetPid() *string
	SetRegionId(v string) *QueryAppTopologyShrinkRequest
	GetRegionId() *string
	SetRpc(v string) *QueryAppTopologyShrinkRequest
	GetRpc() *string
	SetStartTime(v int64) *QueryAppTopologyShrinkRequest
	GetStartTime() *int64
	SetType(v string) *QueryAppTopologyShrinkRequest
	GetType() *string
}

type QueryAppTopologyShrinkRequest struct {
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
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
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

func (s QueryAppTopologyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppTopologyShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAppTopologyShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *QueryAppTopologyShrinkRequest) GetDb() *string {
	return s.Db
}

func (s *QueryAppTopologyShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *QueryAppTopologyShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryAppTopologyShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *QueryAppTopologyShrinkRequest) GetPid() *string {
	return s.Pid
}

func (s *QueryAppTopologyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAppTopologyShrinkRequest) GetRpc() *string {
	return s.Rpc
}

func (s *QueryAppTopologyShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryAppTopologyShrinkRequest) GetType() *string {
	return s.Type
}

func (s *QueryAppTopologyShrinkRequest) SetAppType(v string) *QueryAppTopologyShrinkRequest {
	s.AppType = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetDb(v string) *QueryAppTopologyShrinkRequest {
	s.Db = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetDbName(v string) *QueryAppTopologyShrinkRequest {
	s.DbName = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetEndTime(v int64) *QueryAppTopologyShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetFiltersShrink(v string) *QueryAppTopologyShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetPid(v string) *QueryAppTopologyShrinkRequest {
	s.Pid = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetRegionId(v string) *QueryAppTopologyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetRpc(v string) *QueryAppTopologyShrinkRequest {
	s.Rpc = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetStartTime(v int64) *QueryAppTopologyShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) SetType(v string) *QueryAppTopologyShrinkRequest {
	s.Type = &v
	return s
}

func (s *QueryAppTopologyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
