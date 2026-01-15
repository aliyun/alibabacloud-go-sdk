// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDnsLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplete(v bool) *DescribeInternetDnsLogsResponseBody
	GetComplete() *bool
	SetCurPage(v int32) *DescribeInternetDnsLogsResponseBody
	GetCurPage() *int32
	SetLogs(v *DescribeInternetDnsLogsResponseBodyLogs) *DescribeInternetDnsLogsResponseBody
	GetLogs() *DescribeInternetDnsLogsResponseBodyLogs
	SetPageSize(v int32) *DescribeInternetDnsLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInternetDnsLogsResponseBody
	GetRequestId() *string
	SetTotalPage(v int32) *DescribeInternetDnsLogsResponseBody
	GetTotalPage() *int32
	SetTotalSize(v int32) *DescribeInternetDnsLogsResponseBody
	GetTotalSize() *int32
}

type DescribeInternetDnsLogsResponseBody struct {
	// Indicates whether the log query is precise.
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurPage *int32 `json:"CurPage,omitempty" xml:"CurPage,omitempty"`
	// The queried logs.
	Logs *DescribeInternetDnsLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// Page size for query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identifier.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 5
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// Total quantity.
	//
	// example:
	//
	// 48
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeInternetDnsLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsResponseBody) GetComplete() *bool {
	return s.Complete
}

func (s *DescribeInternetDnsLogsResponseBody) GetCurPage() *int32 {
	return s.CurPage
}

func (s *DescribeInternetDnsLogsResponseBody) GetLogs() *DescribeInternetDnsLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribeInternetDnsLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetDnsLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetDnsLogsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeInternetDnsLogsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeInternetDnsLogsResponseBody) SetComplete(v bool) *DescribeInternetDnsLogsResponseBody {
	s.Complete = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetCurPage(v int32) *DescribeInternetDnsLogsResponseBody {
	s.CurPage = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetLogs(v *DescribeInternetDnsLogsResponseBodyLogs) *DescribeInternetDnsLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetPageSize(v int32) *DescribeInternetDnsLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetRequestId(v string) *DescribeInternetDnsLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetTotalPage(v int32) *DescribeInternetDnsLogsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) SetTotalSize(v int32) *DescribeInternetDnsLogsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBody) Validate() error {
	if s.Logs != nil {
		if err := s.Logs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInternetDnsLogsResponseBodyLogs struct {
	Log []*DescribeInternetDnsLogsResponseBodyLogsLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s DescribeInternetDnsLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsResponseBodyLogs) GetLog() []*DescribeInternetDnsLogsResponseBodyLogsLog {
	return s.Log
}

func (s *DescribeInternetDnsLogsResponseBodyLogs) SetLog(v []*DescribeInternetDnsLogsResponseBodyLogsLog) *DescribeInternetDnsLogsResponseBodyLogs {
	s.Log = v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogs) Validate() error {
	if s.Log != nil {
		for _, item := range s.Log {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInternetDnsLogsResponseBodyLogsLog struct {
	// Parse log ID (can be duplicated).
	//
	// example:
	//
	// 3583
	DnsMsgId *string `json:"DnsMsgId,omitempty" xml:"DnsMsgId,omitempty"`
	Flags    *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// Parse timestamp.
	//
	// example:
	//
	// 1709196249000
	LogTime *int64 `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The protocol type of the domain name resolution query request:
	//
	// - UDP
	//
	// - TCP
	//
	// - HTTP
	//
	// - HTTPS
	//
	// - DOH
	//
	// example:
	//
	// UDP
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	QueryFlags *string `json:"QueryFlags,omitempty" xml:"QueryFlags,omitempty"`
	// The domain name for which you want to query Domain Name System (DNS) records.
	//
	// example:
	//
	// example.com
	QueryName *string `json:"QueryName,omitempty" xml:"QueryName,omitempty"`
	// Record type.
	//
	// example:
	//
	// A
	QueryType         *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResponseTimestamp *string `json:"ResponseTimestamp,omitempty" xml:"ResponseTimestamp,omitempty"`
	// Parse response time.
	//
	// example:
	//
	// 0
	Rt *int32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// Parse server IP.
	//
	// example:
	//
	// 140.205.XX.XX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// Source IP address.
	//
	// example:
	//
	// 59.82.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Response status.
	//
	// example:
	//
	// NOERROR
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The value set for the edns-client-subnet option.
	//
	// example:
	//
	// 170.33.XX.XX
	SubnetIp *string `json:"SubnetIp,omitempty" xml:"SubnetIp,omitempty"`
	// Array of parsing results.
	Value *DescribeInternetDnsLogsResponseBodyLogsLogValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
	// The zone name.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInternetDnsLogsResponseBodyLogsLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsResponseBodyLogsLog) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetDnsMsgId() *string {
	return s.DnsMsgId
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetFlags() *string {
	return s.Flags
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetLogTime() *int64 {
	return s.LogTime
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetQueryFlags() *string {
	return s.QueryFlags
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetQueryName() *string {
	return s.QueryName
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetResponseTimestamp() *string {
	return s.ResponseTimestamp
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetRt() *int32 {
	return s.Rt
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetServerIp() *string {
	return s.ServerIp
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetStatus() *string {
	return s.Status
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetSubnetIp() *string {
	return s.SubnetIp
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetValue() *DescribeInternetDnsLogsResponseBodyLogsLogValue {
	return s.Value
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetDnsMsgId(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.DnsMsgId = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetFlags(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.Flags = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetLogTime(v int64) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.LogTime = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetProtocol(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.Protocol = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetQueryFlags(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.QueryFlags = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetQueryName(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.QueryName = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetQueryType(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.QueryType = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetResponseTimestamp(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.ResponseTimestamp = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetRt(v int32) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.Rt = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetServerIp(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.ServerIp = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetSourceIp(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetStatus(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.Status = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetSubnetIp(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.SubnetIp = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetValue(v *DescribeInternetDnsLogsResponseBodyLogsLogValue) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.Value = v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) SetZoneName(v string) *DescribeInternetDnsLogsResponseBodyLogsLog {
	s.ZoneName = &v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLog) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInternetDnsLogsResponseBodyLogsLogValue struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeInternetDnsLogsResponseBodyLogsLogValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsResponseBodyLogsLogValue) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLogValue) GetValue() []*string {
	return s.Value
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLogValue) SetValue(v []*string) *DescribeInternetDnsLogsResponseBodyLogsLogValue {
	s.Value = v
	return s
}

func (s *DescribeInternetDnsLogsResponseBodyLogsLogValue) Validate() error {
	return dara.Validate(s)
}
