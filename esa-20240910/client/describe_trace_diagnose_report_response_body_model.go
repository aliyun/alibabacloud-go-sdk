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
	// IP address of the local DNS server.
	//
	// example:
	//
	// 47.xx.112.120
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// Client information.
	ClientInfo *DescribeTraceDiagnoseReportResponseBodyClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	// Client IP.
	//
	// example:
	//
	// 33.7.98.136
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// Creation time. Format: yyyy-MM-dd HH:mm:ss, timezone: +08:00.
	//
	// example:
	//
	// 2024-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Diagnostic ID.
	//
	// example:
	//
	// vpndgn-hn316ixao7ut50ybl5qui
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// Diagnostic report link.
	//
	// example:
	//
	// https://cdnskyeye.alibaba-inc.com/#/databrain/detectionPlatformAction?module=trace&id=xxx&instance_id=xxxx&type=copy&edition_id=xxxxx
	DiagnoseReportLink *string `json:"DiagnoseReportLink,omitempty" xml:"DiagnoseReportLink,omitempty"`
	// Diagnostic link.
	//
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose/xxxxxx
	DiagnoseUrl *string `json:"DiagnoseUrl,omitempty" xml:"DiagnoseUrl,omitempty"`
	// The diagnosed domain name.
	//
	// example:
	//
	// www.pecmnr.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Expiration time. Timestamp in seconds.
	//
	// example:
	//
	// 1678701915
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Remaining available diagnosis count.
	//
	// example:
	//
	// 10
	RemainDiagnoseTimes *int64 `json:"RemainDiagnoseTimes,omitempty" xml:"RemainDiagnoseTimes,omitempty"`
	// Diagnostic report details.
	Report *DescribeTraceDiagnoseReportResponseBodyReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Report generation status.
	//
	// 0: Success.
	//
	// 1: Failure.
	//
	// 2: Timeout.
	//
	// 3: Running.
	//
	// 4: Waiting.
	//
	// example:
	//
	// ok
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Status of the diagnostic link.
	//
	// 1: Active.
	//
	// 0: Expired.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task ID.
	//
	// example:
	//
	// xxxxxxx-xxxxx-xxxxx-xxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Time consumed for report generation, in seconds.
	//
	// example:
	//
	// 10
	TimeConsuming *int64 `json:"TimeConsuming,omitempty" xml:"TimeConsuming,omitempty"`
	// Trace display link.
	//
	// example:
	//
	// https://tracing-sk.alibaba-inc.com/trace/xxxxxxxxxxxxxx
	TraceDisplayLink *string `json:"TraceDisplayLink,omitempty" xml:"TraceDisplayLink,omitempty"`
	// Diagnostic trace ID.
	//
	// example:
	//
	// 0000006xxxxxxxxxxxx533427e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
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
	// Browser.
	//
	// example:
	//
	// Chrome
	BrowserInfo *string `json:"BrowserInfo,omitempty" xml:"BrowserInfo,omitempty"`
	// Operating system name.
	//
	// example:
	//
	// Macintosh
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// Version.
	//
	// example:
	//
	// Mozilla/5.0 (Macintosh Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36
	UaString *string `json:"UaString,omitempty" xml:"UaString,omitempty"`
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
	// Client information.
	//
	// example:
	//
	// Client IP : xxx.xxx.xxx.xxx DNS_IP : xxx.xxx.xxx.xxx
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	// Diagnostic result.
	//
	// example:
	//
	// Diagnosis Result: Normal Scheduling Result: Normal Check Item Details: Exception Item: [Request Normal]
	DiagnoseResult *string `json:"DiagnoseResult,omitempty" xml:"DiagnoseResult,omitempty"`
	// Client request response header.
	//
	// example:
	//
	// <br>connection: keep-alive<br>content-type: application/json
	ResponseHeader *string `json:"ResponseHeader,omitempty" xml:"ResponseHeader,omitempty"`
	// Static snapshot page.
	//
	// example:
	//
	// https://cdn.dns-detect.alicdn.com/diagnose_v2#/snapshot?data=H4sIAAAAAAACA%2B1aeW%2FcxhX%2FKgSLFkm7pGbI4bWGUSiSEqmxZMG7Tow4hjAkh1pGFLkmuVorgYE4QeI0cc7eB3IgLpIWTdO4ORzXQb6MV5b%2F6lfomyG5u1xpbUW2mwSNpAWWM2%2FmHfPm935DzTNyHm6yLKebXbkpa0gzFWQo2Gkjp4nNpm6qBJtPyA052WLpVsj6cvMZ2Y%2BzeZYzLw%2BTmD%2Fn7FwOo%2BdXWjsfvT%2B4ehXEvSRKUmj7kaF5BFP5fEOOEo9GJ9OoDerGRt18%2BfOdZy%2FcfPcCQWRw%2FbmJ4UFAfBLw4Vm42YtoziaG73z43s4r7%2Bx89srux7%2B59fyHg5devN0kfrJJw3iFbjJoH3xxZXD54s57L8lVxxxL8zAIPVDDp66cSVmWRFtsrsO8jTHVS6tb5u4H7%2B%2B8%2FcbtVAq355I45vHaCvPtceNfemPw8js3X7t44%2Bo1jNCP%2F3P90u7Xb9969o87z34wPZTcD5bTMMr4VGF3KQ4S%2Fk1oWuLriJGpYqzqRMUazBDBgvFWW1MdDZoM1YTWsHu8y1Kai8lv%2Ff7rWxcvDV584eZXb4q%2B1TTZCmNPxOnS725c%2B%2Fvg6nOiYy7pxXkKXsg3rn40%2BNNXPHhxtmeqG9feLHqmTAQ9%2B8wEiZj3wNY87bGJ2MnN08%2FIvTQC8X6%2Fr54NadxnseolmyJGPijQEKpPAU89z2NZoUpu6g15k%2Fkwcr4H1ors1bHWkCkkN11nY60azOSXjzDZac2wVQcVP5qmG1oDBqq6U%2FzYBnKchm7bKi5lMHEQPnO%2BMW4yjcLt3t1bjG17H4uxY9Ytxoah6qU1yCIGMhvYtlRz3ECtoRFDtSq%2FHBtp%2BoTRZ8%2FevcGOsY%2FBGiYTBuuQm2VAHc0gDmlgh6j2WIzBcx0T1axJTRjs0tDv3b3NJrH2i7JorlltErBxLM4NaDHHTeQtBqqtxZnzZxqym4iUTkFzCvld2sLzLqTrcZKF2XKeLvngFnjSah0r90E75OCFLTCaZhkHrCQ%2BwbJelGeFaDeM1%2BcFPhTPi%2B32am2s6cA%2BzjhOENMAAZamSTonQgUY3JBzmq6zvNq3cyvFNDBCfC2R4otP4a%2FYsWneTmmcBSxtFWEFKR7mCjl7Rewq21oszld6m0UDTRld4NMu9uhDLJS59mxjIfbbo7KELcsB7MK2AX9j4VmFyar4FDYvZd2F0tx2ktOo8FfXYYtnSS%2F12FJ3yyz6AX9aUNFYKuAS8Y2rQv7psEvGnuzSoJbwch%2BTiG3oPM8Ekon4DLGsWwZQWLTCYHC8vt4TIWPd2oqVw4XgXCeMaW0VH07STZrnrPR0kVH%2Bja%2FqDFaxBEsmrSS59DBM4j%2BZPhnDWucQYuUY6Ms7TQnxxhbLlbkk2QhZU6Jefy33jmqB5WLNZaPwIqLbJrYZ1glyYKPZpuXqtmH5gYlp4Gg6CVwMAHKkS%2FPO0Zkji3nePR5H20eW6Tlldp0d1aFzUpvnx2sZ8%2B6zxnko2U2pDXtbwo60TLclzmWkgsVoRHpkuc3FYAkhrspsFCV95XgarodxU%2FppYTJPhqa00Jrlj4%2FRdLspzUI6d3NlIQYggXGiI6RNicWZ4lGvw3SsRhpLsGkq2DoNZbeBGvOrZxrTJewG%2F8CiKVizUWN5XNZyarIYNdB4N7FUL7aww7usPYrqvRibU5QQbSRngwLu1CmlFeZMmeP9SksgZVNaXmq1eOdsFCqtfhjkyiNR4tJIaQEsctbYlEYLWc4ixEb909bDqNajGrLAIagpJbAiEmBGF7CVSeTUKUlg07iksLHUztsX6HrEQr8pTU8uLsY%2FsHmAop6g%2FXHImANGVmGGB1WkjrI64E2a9HL2i6yCsM1zWW382NYX%2FKjq5ZBIBUUWGBV4nmlBiVM8x7AU4miB4riupvgAOQwh6um%2BXZqwmrK8BNQK7IeoLOgd0AyolZpmqhrHb2FgobVCuVIIQExXsSVQ7Fw%2BsgxEBGesgHAStnOvuzcOXhWpx4GHx16HuhzM%2BmGaM45J%2B6A%2Fx6j5pB9HCfVbXcYRDBXIthzGxczyhFjRiorGE2UmjNWmKjkeSvztUQUR0R%2B8%2BhoQTLkwlBc3XRS3c93qoFIVIMCOkfJ6sasyAQKwx5s4KwOY9ZIpQQ7L%2BgNbxqUurXxbaZVhRMZYpRrlnKgu9SXXdTyqIIV7L76w%2B9YfBq9f3r34t50r%2Fx5cula6c4J5LNxi%2FrCm5vuVTsRJEKx1m69%2BUX0K6Yp%2FDCM%2F2heFQD8FbDjBzvZg0j1BAU0ZMCRe4aoaBK094CUi67Hp6Lph246t2Zhgwh2O4fRVRjKA8tZL2Qla5W%2BxWnvTohU%2BXTNurGDK5TFhvjjj3LaUQ0Bd1t7ucm2h78mjiCzT3OsUs5WuFjkNTatAATcB9OOkAFBRYwVS8kqbJlG96yT4zktTnDeluU6abLIZwxI4KooJlJsZUXAWkyxv%2BjTrZF7SZSrojMNefwPAeQyveCoXrOLLS4MX%2Fjn4%2BJ3B1eeLAxSnJSLGvqEFPtKJQm3bVQgzdMUB%2Bqn4VIeou55p0IBnbRZNrB7nzYdknwY5NPu0ze8p%2BySGrh%2BCfRIL3z%2F2iUz8HWOf0%2Br0fgSUWo5j4FHFJpquEThpMwsR2zKohjUfiqbNCDVNm3iIBpDPcBQ8NAG9fxoPSEDr9BG4GBC%2BnmMoGAE3xHtJXU3AAF5nTON1GqmJOnXuiDEAi20QHQ6sxh41o17NQA3%2B2V8HRiM5POKOB6RwExT7rkkn%2BWak8wCHgBGZnJooU09YU08Mh6SeyLpX1NOAo5QeoEDxLYKgQECpcOzAVwLPtUzX1d0AW1Opp1GnnoS%2FP4SNTxzVwVO5p%2B4AhhFVQ5ZqobsmnzwSQ%2FK52OuwDk1y9p3mnqZl3nfuORnkO5BP0z4o9xQV6x5yT%2FID9zws97wHPHM%2F%2BvrNuOfgzeuDty7tfPbXm7%2F88tvnnpp2aOr5fX3vSbCDD8E8%2BXug%2B8c8%2BTuJOzPPReYK776Nd551rmVawI%2B82NI157RO9lCgiW6g%2BvtTIBOPyRl1kqVxGmWZhnXa1vcoqHXa2rS3c%2BZQikwQrG%2FAcO74ssxyDEenLhlbT8cghqaxKVx9%2BgCsa4hgz3Cp5wfUdhAUJNMknu8xIBfOoV8W30eNt%2BfqehXFA3PQg747vi3jPTiPns5fD0k1yb1imh6jrufbQDI94ijE1XyFEuIqvh5Yhm76xKZ0GtPkuD5GNOE0BNQGEAkbqkZuxzQdWwU2ajh3zTPJGM18okPj9adCupF8t19yGkj%2FXxDNsRjf6R2nflCaKcpTjWbuXPkXOHfzzxd%2BYJf%2FR%2Bzy%2BruDf3w%2BeP39e0Ytz9Tv%2BRT%2F8YdwPMbSMNgG8Kf8tkzW63aTNF9j4nJMQKOMS4Xra%2BVtn%2BHAoqnIqrmV2eWFpvRAnMTsQQiGBCmapH7WfDKWpPF%2Fy%2FBnYsFZWVeh8GtYm2jQUb0BO059imLE6NkCbfAzVDi0QezVwpWwu0VGNxyGbebQPY9xh891YVv4E41BIO4nsbW8QI3hVTAiIauJSFPD0s%2F49QXpZHtOrk1VH2IDq6mGoL1DqnsYcpyMTDnfkDtQrIeJOfjVq4Nrvx5c%2FvTG1%2B%2FtXPh497UvBq%2F%2FdnD5k91P%2FyIXouO0HbQfd58S16oKQ0xsq5q0mXGdPClLIMQqKRq5GDDMagsC0ogNbVSD8hGHHjWDjUlWXbVKuAfL1AvjPMk6MMRNk35WbCqxWzhUUC6TPB1GEZ0xVCQ9MBxwRFoCvhoB9fCk4y3plITRGjbWrAel2W43Yo8z99EwnzF0S9VN6YFHF9vLxxpSFG4w6RHmbSQPVlsSE1tF%2FFdq0YCmYTlE0GtOqBZi6kZ8qYucEGjSyhN%2Bp2WiK4mPhTEb5sQwHUpIIevijhzEK4w3%2BCsZQNecn9kgKjlAyklx%2FYYvTNacmdl378%2F8nEbhGid4PnN762t8prXQP1oAuWHZWEHw61gYeB0KGMYscC3MgNthD3m6ZzKNIvMnfCE4Cgtw%2BtYifP6%2Fq6Vk7DkpAAA%3D
	StaticHtml *string `json:"StaticHtml,omitempty" xml:"StaticHtml,omitempty"`
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
