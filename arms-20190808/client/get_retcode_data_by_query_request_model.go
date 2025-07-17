// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeDataByQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetRetcodeDataByQueryRequest
	GetFrom() *int64
	SetPid(v string) *GetRetcodeDataByQueryRequest
	GetPid() *string
	SetQuery(v string) *GetRetcodeDataByQueryRequest
	GetQuery() *string
	SetRegionId(v string) *GetRetcodeDataByQueryRequest
	GetRegionId() *string
	SetTo(v int64) *GetRetcodeDataByQueryRequest
	GetTo() *int64
}

type GetRetcodeDataByQueryRequest struct {
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1668687302
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the application.
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of an application. The URL in the address bar contains the process ID (PID) of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is eb4zdose6v%409781be0f44d\\*\\*\\*\\*, you must replace %40 with an at sign (@) to obtain eb4zdose6v@9781be0f44d\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The query statement that conforms to the query syntax of a Log Service Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// t : pv|select sum(times) as pv , approx_distinct(uid) as uv , (date-date%3600000) as date  group by date
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1668688000
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetRetcodeDataByQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeDataByQueryRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeDataByQueryRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetRetcodeDataByQueryRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeDataByQueryRequest) GetQuery() *string {
	return s.Query
}

func (s *GetRetcodeDataByQueryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRetcodeDataByQueryRequest) GetTo() *int64 {
	return s.To
}

func (s *GetRetcodeDataByQueryRequest) SetFrom(v int64) *GetRetcodeDataByQueryRequest {
	s.From = &v
	return s
}

func (s *GetRetcodeDataByQueryRequest) SetPid(v string) *GetRetcodeDataByQueryRequest {
	s.Pid = &v
	return s
}

func (s *GetRetcodeDataByQueryRequest) SetQuery(v string) *GetRetcodeDataByQueryRequest {
	s.Query = &v
	return s
}

func (s *GetRetcodeDataByQueryRequest) SetRegionId(v string) *GetRetcodeDataByQueryRequest {
	s.RegionId = &v
	return s
}

func (s *GetRetcodeDataByQueryRequest) SetTo(v int64) *GetRetcodeDataByQueryRequest {
	s.To = &v
	return s
}

func (s *GetRetcodeDataByQueryRequest) Validate() error {
	return dara.Validate(s)
}
