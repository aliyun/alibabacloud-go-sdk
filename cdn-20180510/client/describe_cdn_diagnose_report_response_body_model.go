// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnDiagnoseReportResponseBodyContent) *DescribeCdnDiagnoseReportResponseBody
	GetContent() *DescribeCdnDiagnoseReportResponseBodyContent
	SetRequestId(v string) *DescribeCdnDiagnoseReportResponseBody
	GetRequestId() *string
}

type DescribeCdnDiagnoseReportResponseBody struct {
	Content *DescribeCdnDiagnoseReportResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 64D28B53-xxxxx-xxx-xxxxx-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponseBody) GetContent() *DescribeCdnDiagnoseReportResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDiagnoseReportResponseBody) SetContent(v *DescribeCdnDiagnoseReportResponseBodyContent) *DescribeCdnDiagnoseReportResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBody) SetRequestId(v string) *DescribeCdnDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDiagnoseReportResponseBodyContent struct {
	// example:
	//
	// 1077********7468
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// example:
	//
	// xx.xx.xx.xx
	ClientAddr *string                                                 `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientInfo *DescribeCdnDiagnoseReportResponseBodyContentClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	// example:
	//
	// xx.xx.xx.xx
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 2022-12-10 15:11:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// f2xxxxd5
	DiagnoseId         *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	DiagnoseReportLink *string `json:"DiagnoseReportLink,omitempty" xml:"DiagnoseReportLink,omitempty"`
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose/xxxxxx
	DiagnoseUrl *string `json:"DiagnoseUrl,omitempty" xml:"DiagnoseUrl,omitempty"`
	// example:
	//
	// http://www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1678701915
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 10
	RemainDiagnoseTimes *int64                                              `json:"RemainDiagnoseTimes,omitempty" xml:"RemainDiagnoseTimes,omitempty"`
	Report              *DescribeCdnDiagnoseReportResponseBodyContentReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xxxxxxx-xxxxx-xxxxx-xxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1s
	TimeConsuming *int64 `json:"TimeConsuming,omitempty" xml:"TimeConsuming,omitempty"`
	// example:
	//
	// https://tracing-sk.alibaba-inc.com/trace/xxxxxxxxxxxxxx
	TraceDisplayLink *string `json:"TraceDisplayLink,omitempty" xml:"TraceDisplayLink,omitempty"`
	// example:
	//
	// 0000006xxxxxxxxxxxx533427e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeCdnDiagnoseReportResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetAliuid() *string {
	return s.Aliuid
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetClientInfo() *DescribeCdnDiagnoseReportResponseBodyContentClientInfo {
	return s.ClientInfo
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetDiagnoseReportLink() *string {
	return s.DiagnoseReportLink
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetDiagnoseUrl() *string {
	return s.DiagnoseUrl
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetRemainDiagnoseTimes() *int64 {
	return s.RemainDiagnoseTimes
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetReport() *DescribeCdnDiagnoseReportResponseBodyContentReport {
	return s.Report
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetState() *string {
	return s.State
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetTimeConsuming() *int64 {
	return s.TimeConsuming
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetTraceDisplayLink() *string {
	return s.TraceDisplayLink
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetAliuid(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.Aliuid = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetClientAddr(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.ClientAddr = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetClientInfo(v *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.ClientInfo = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetClientIp(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.ClientIp = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetCreateTime(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetDiagnoseId(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.DiagnoseId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetDiagnoseReportLink(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.DiagnoseReportLink = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetDiagnoseUrl(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.DiagnoseUrl = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetDomain(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.Domain = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetExpireTime(v int64) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetRemainDiagnoseTimes(v int64) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.RemainDiagnoseTimes = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetReport(v *DescribeCdnDiagnoseReportResponseBodyContentReport) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.Report = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetState(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.State = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetStatus(v int64) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.Status = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetTaskId(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.TaskId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetTimeConsuming(v int64) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.TimeConsuming = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetTraceDisplayLink(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.TraceDisplayLink = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) SetTraceId(v string) *DescribeCdnDiagnoseReportResponseBodyContent {
	s.TraceId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContent) Validate() error {
	if s.ClientInfo != nil {
		if err := s.ClientInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDiagnoseReportResponseBodyContentClientInfo struct {
	// example:
	//
	// Chrome
	BrowserInfo *string `json:"BrowserInfo,omitempty" xml:"BrowserInfo,omitempty"`
	// example:
	//
	// Macintosh
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// Mozilla/5.0 (Macintosh Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36
	UAString *string `json:"UAString,omitempty" xml:"UAString,omitempty"`
}

func (s DescribeCdnDiagnoseReportResponseBodyContentClientInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponseBodyContentClientInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) GetBrowserInfo() *string {
	return s.BrowserInfo
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) GetOs() *string {
	return s.Os
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) GetUAString() *string {
	return s.UAString
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) SetBrowserInfo(v string) *DescribeCdnDiagnoseReportResponseBodyContentClientInfo {
	s.BrowserInfo = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) SetOs(v string) *DescribeCdnDiagnoseReportResponseBodyContentClientInfo {
	s.Os = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) SetUAString(v string) *DescribeCdnDiagnoseReportResponseBodyContentClientInfo {
	s.UAString = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentClientInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDiagnoseReportResponseBodyContentReport struct {
	ClientInfo     *string                                                      `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	DiagnoseResult *string                                                      `json:"DiagnoseResult,omitempty" xml:"DiagnoseResult,omitempty"`
	L1Node         *string                                                      `json:"L1Node,omitempty" xml:"L1Node,omitempty"`
	L1Tengine      *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine `json:"L1Tengine,omitempty" xml:"L1Tengine,omitempty" type:"Struct"`
	SourceInfo     []*string                                                    `json:"SourceInfo,omitempty" xml:"SourceInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnDiagnoseReportResponseBodyContentReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponseBodyContentReport) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) GetClientInfo() *string {
	return s.ClientInfo
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) GetDiagnoseResult() *string {
	return s.DiagnoseResult
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) GetL1Node() *string {
	return s.L1Node
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) GetL1Tengine() *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	return s.L1Tengine
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) GetSourceInfo() []*string {
	return s.SourceInfo
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) SetClientInfo(v string) *DescribeCdnDiagnoseReportResponseBodyContentReport {
	s.ClientInfo = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) SetDiagnoseResult(v string) *DescribeCdnDiagnoseReportResponseBodyContentReport {
	s.DiagnoseResult = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) SetL1Node(v string) *DescribeCdnDiagnoseReportResponseBodyContentReport {
	s.L1Node = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) SetL1Tengine(v *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) *DescribeCdnDiagnoseReportResponseBodyContentReport {
	s.L1Tengine = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) SetSourceInfo(v []*string) *DescribeCdnDiagnoseReportResponseBodyContentReport {
	s.SourceInfo = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReport) Validate() error {
	if s.L1Tengine != nil {
		if err := s.L1Tengine.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine struct {
	// example:
	//
	// 150.xxx.xxx.197:81
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Unmatched CNAME.
	ErrorLog *string `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty"`
	// example:
	//
	// 107ms
	Fbt   *string `json:"Fbt,omitempty" xml:"Fbt,omitempty"`
	Intro *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	// example:
	//
	// aG9zd******8tbWVtYmV
	ReqHeader *string `json:"ReqHeader,omitempty" xml:"ReqHeader,omitempty"`
	// example:
	//
	// 2023-02-06 12:17:47
	ReqTime *string `json:"ReqTime,omitempty" xml:"ReqTime,omitempty"`
	// example:
	//
	// U2VydmV*******Qo
	RespHeader *string `json:"RespHeader,omitempty" xml:"RespHeader,omitempty"`
	// example:
	//
	// 2892
	RespSize *int64 `json:"RespSize,omitempty" xml:"RespSize,omitempty"`
	// example:
	//
	// 107ms
	Rt *string `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// example:
	//
	// cn4461
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
}

func (s DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetAddr() *string {
	return s.Addr
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetCode() *string {
	return s.Code
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetErrorLog() *string {
	return s.ErrorLog
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetFbt() *string {
	return s.Fbt
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetIntro() *string {
	return s.Intro
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetReqHeader() *string {
	return s.ReqHeader
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetReqTime() *string {
	return s.ReqTime
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetRespHeader() *string {
	return s.RespHeader
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetRespSize() *int64 {
	return s.RespSize
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetRt() *string {
	return s.Rt
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) GetStation() *string {
	return s.Station
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetAddr(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Addr = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetCode(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Code = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetErrorLog(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.ErrorLog = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetFbt(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Fbt = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetIntro(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Intro = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetReqHeader(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.ReqHeader = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetReqTime(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.ReqTime = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetRespHeader(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.RespHeader = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetRespSize(v int64) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.RespSize = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetRt(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Rt = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) SetStation(v string) *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine {
	s.Station = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponseBodyContentReportL1Tengine) Validate() error {
	return dara.Validate(s)
}
