// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientAddr(v string) *DescribeTraceDiagnoseReportResponseBody
	GetClientAddr() *string
	SetClientInfo(v *DescribeTraceDiagnoseReportResponseBodyClientInfo) *DescribeTraceDiagnoseReportResponseBody
	GetClientInfo() *DescribeTraceDiagnoseReportResponseBodyClientInfo
	SetClientIp(v string) *DescribeTraceDiagnoseReportResponseBody
	GetClientIp() *string
	SetCreateTime(v string) *DescribeTraceDiagnoseReportResponseBody
	GetCreateTime() *string
	SetDiagnoseId(v string) *DescribeTraceDiagnoseReportResponseBody
	GetDiagnoseId() *string
	SetDiagnoseReportLink(v string) *DescribeTraceDiagnoseReportResponseBody
	GetDiagnoseReportLink() *string
	SetDiagnoseUrl(v string) *DescribeTraceDiagnoseReportResponseBody
	GetDiagnoseUrl() *string
	SetDomain(v string) *DescribeTraceDiagnoseReportResponseBody
	GetDomain() *string
	SetExpireTime(v int64) *DescribeTraceDiagnoseReportResponseBody
	GetExpireTime() *int64
	SetRemainDiagnoseTimes(v int64) *DescribeTraceDiagnoseReportResponseBody
	GetRemainDiagnoseTimes() *int64
	SetReport(v *DescribeTraceDiagnoseReportResponseBodyReport) *DescribeTraceDiagnoseReportResponseBody
	GetReport() *DescribeTraceDiagnoseReportResponseBodyReport
	SetRequestId(v string) *DescribeTraceDiagnoseReportResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeTraceDiagnoseReportResponseBody
	GetState() *string
	SetStatus(v int64) *DescribeTraceDiagnoseReportResponseBody
	GetStatus() *int64
	SetTaskId(v string) *DescribeTraceDiagnoseReportResponseBody
	GetTaskId() *string
	SetTimeConsuming(v int64) *DescribeTraceDiagnoseReportResponseBody
	GetTimeConsuming() *int64
	SetTraceDisplayLink(v string) *DescribeTraceDiagnoseReportResponseBody
	GetTraceDisplayLink() *string
	SetTraceId(v string) *DescribeTraceDiagnoseReportResponseBody
	GetTraceId() *string
}

type DescribeTraceDiagnoseReportResponseBody struct {
	ClientAddr          *string                                            `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientInfo          *DescribeTraceDiagnoseReportResponseBodyClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	ClientIp            *string                                            `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	CreateTime          *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DiagnoseId          *string                                            `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	DiagnoseReportLink  *string                                            `json:"DiagnoseReportLink,omitempty" xml:"DiagnoseReportLink,omitempty"`
	DiagnoseUrl         *string                                            `json:"DiagnoseUrl,omitempty" xml:"DiagnoseUrl,omitempty"`
	Domain              *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExpireTime          *int64                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	RemainDiagnoseTimes *int64                                             `json:"RemainDiagnoseTimes,omitempty" xml:"RemainDiagnoseTimes,omitempty"`
	Report              *DescribeTraceDiagnoseReportResponseBodyReport     `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State               *string                                            `json:"State,omitempty" xml:"State,omitempty"`
	Status              *int64                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId              *string                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TimeConsuming       *int64                                             `json:"TimeConsuming,omitempty" xml:"TimeConsuming,omitempty"`
	TraceDisplayLink    *string                                            `json:"TraceDisplayLink,omitempty" xml:"TraceDisplayLink,omitempty"`
	TraceId             *string                                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeTraceDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetClientInfo() *DescribeTraceDiagnoseReportResponseBodyClientInfo {
	return s.ClientInfo
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetDiagnoseReportLink() *string {
	return s.DiagnoseReportLink
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetDiagnoseUrl() *string {
	return s.DiagnoseUrl
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetRemainDiagnoseTimes() *int64 {
	return s.RemainDiagnoseTimes
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetReport() *DescribeTraceDiagnoseReportResponseBodyReport {
	return s.Report
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetTimeConsuming() *int64 {
	return s.TimeConsuming
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetTraceDisplayLink() *string {
	return s.TraceDisplayLink
}

func (s *DescribeTraceDiagnoseReportResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetClientAddr(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.ClientAddr = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetClientInfo(v *DescribeTraceDiagnoseReportResponseBodyClientInfo) *DescribeTraceDiagnoseReportResponseBody {
	s.ClientInfo = v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetClientIp(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.ClientIp = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetCreateTime(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetDiagnoseId(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.DiagnoseId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetDiagnoseReportLink(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.DiagnoseReportLink = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetDiagnoseUrl(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.DiagnoseUrl = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetDomain(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetExpireTime(v int64) *DescribeTraceDiagnoseReportResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetRemainDiagnoseTimes(v int64) *DescribeTraceDiagnoseReportResponseBody {
	s.RemainDiagnoseTimes = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetReport(v *DescribeTraceDiagnoseReportResponseBodyReport) *DescribeTraceDiagnoseReportResponseBody {
	s.Report = v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetRequestId(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetState(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.State = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetStatus(v int64) *DescribeTraceDiagnoseReportResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetTaskId(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetTimeConsuming(v int64) *DescribeTraceDiagnoseReportResponseBody {
	s.TimeConsuming = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetTraceDisplayLink(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.TraceDisplayLink = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) SetTraceId(v string) *DescribeTraceDiagnoseReportResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBody) Validate() error {
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

type DescribeTraceDiagnoseReportResponseBodyClientInfo struct {
	BrowserInfo *string `json:"BrowserInfo,omitempty" xml:"BrowserInfo,omitempty"`
	Os          *string `json:"Os,omitempty" xml:"Os,omitempty"`
	UaString    *string `json:"UaString,omitempty" xml:"UaString,omitempty"`
}

func (s DescribeTraceDiagnoseReportResponseBodyClientInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceDiagnoseReportResponseBodyClientInfo) GoString() string {
	return s.String()
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) GetBrowserInfo() *string {
	return s.BrowserInfo
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) GetOs() *string {
	return s.Os
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) GetUaString() *string {
	return s.UaString
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) SetBrowserInfo(v string) *DescribeTraceDiagnoseReportResponseBodyClientInfo {
	s.BrowserInfo = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) SetOs(v string) *DescribeTraceDiagnoseReportResponseBodyClientInfo {
	s.Os = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) SetUaString(v string) *DescribeTraceDiagnoseReportResponseBodyClientInfo {
	s.UaString = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyClientInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceDiagnoseReportResponseBodyReport struct {
	ClientInfo     *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	DiagnoseResult *string `json:"DiagnoseResult,omitempty" xml:"DiagnoseResult,omitempty"`
	ResponseHeader *string `json:"ResponseHeader,omitempty" xml:"ResponseHeader,omitempty"`
	StaticHtml     *string `json:"StaticHtml,omitempty" xml:"StaticHtml,omitempty"`
}

func (s DescribeTraceDiagnoseReportResponseBodyReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceDiagnoseReportResponseBodyReport) GoString() string {
	return s.String()
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) GetClientInfo() *string {
	return s.ClientInfo
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) GetDiagnoseResult() *string {
	return s.DiagnoseResult
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) GetResponseHeader() *string {
	return s.ResponseHeader
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) GetStaticHtml() *string {
	return s.StaticHtml
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) SetClientInfo(v string) *DescribeTraceDiagnoseReportResponseBodyReport {
	s.ClientInfo = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) SetDiagnoseResult(v string) *DescribeTraceDiagnoseReportResponseBodyReport {
	s.DiagnoseResult = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) SetResponseHeader(v string) *DescribeTraceDiagnoseReportResponseBodyReport {
	s.ResponseHeader = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) SetStaticHtml(v string) *DescribeTraceDiagnoseReportResponseBodyReport {
	s.StaticHtml = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponseBodyReport) Validate() error {
	return dara.Validate(s)
}
