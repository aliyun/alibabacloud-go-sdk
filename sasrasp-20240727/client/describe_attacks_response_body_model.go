// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttacks(v []*DescribeAttacksResponseBodyAttacks) *DescribeAttacksResponseBody
	GetAttacks() []*DescribeAttacksResponseBodyAttacks
	SetRequestId(v string) *DescribeAttacksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAttacksResponseBody
	GetTotalCount() *int64
}

type DescribeAttacksResponseBody struct {
	Attacks []*DescribeAttacksResponseBodyAttacks `json:"Attacks,omitempty" xml:"Attacks,omitempty" type:"Repeated"`
	// example:
	//
	// 4E7772EA-1AD0-5583-BA55-A4E7B65F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 289
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAttacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttacksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttacksResponseBody) GetAttacks() []*DescribeAttacksResponseBodyAttacks {
	return s.Attacks
}

func (s *DescribeAttacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAttacksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAttacksResponseBody) SetAttacks(v []*DescribeAttacksResponseBodyAttacks) *DescribeAttacksResponseBody {
	s.Attacks = v
	return s
}

func (s *DescribeAttacksResponseBody) SetRequestId(v string) *DescribeAttacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttacksResponseBody) SetTotalCount(v int64) *DescribeAttacksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAttacksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAttacksResponseBodyAttacks struct {
	// example:
	//
	// /app
	AppDir *string `json:"AppDir,omitempty" xml:"AppDir,omitempty"`
	// example:
	//
	// 661dece60bb1
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// zhxy
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// null
	Avd *string `json:"Avd,omitempty" xml:"Avd,omitempty"`
	// example:
	//
	// jar
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// example:
	//
	// high
	Confidence *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// {\\"Exist\\": \\"false\\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1024
	ContentLength *int64 `json:"ContentLength,omitempty" xml:"ContentLength,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// {\\"All\\": 12, \\"Online\\": 9}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {\\"X-Total-Count\\": 1}
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// 100.116.60.78
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// hostnames
	Hostname           *string                                                 `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InputParamItemList []*DescribeAttacksResponseBodyAttacksInputParamItemList `json:"InputParamItemList,omitempty" xml:"InputParamItemList,omitempty" type:"Repeated"`
	InstallType        *int64                                                  `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	// example:
	//
	// 47.92.92.85
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 1.8.0_342
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1.0
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// example:
	//
	// 1727281449756
	MaxTimestamp *int64  `json:"MaxTimestamp,omitempty" xml:"MaxTimestamp,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// null
	MiddlewareInstanceId *string `json:"MiddlewareInstanceId,omitempty" xml:"MiddlewareInstanceId,omitempty"`
	// example:
	//
	// 1727281449756
	MinTimestamp *int64 `json:"MinTimestamp,omitempty" xml:"MinTimestamp,omitempty"`
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// linux
	OsArch *string `json:"OsArch,omitempty" xml:"OsArch,omitempty"`
	// example:
	//
	// aliyun_V3_5.1
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// example:
	//
	// PUBLIC_CLOUD
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// example:
	//
	// 245256
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// example:
	//
	// 100
	PayloadLength *int64 `json:"PayloadLength,omitempty" xml:"PayloadLength,omitempty"`
	// example:
	//
	// 29719
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// php
	RaspVersion *string `json:"RaspVersion,omitempty" xml:"RaspVersion,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 192.168.1.1
	Remote *string `json:"Remote,omitempty" xml:"Remote,omitempty"`
	// example:
	//
	// [{\\"TagKey\\": \\"acs:rm:rgId\\", \\"TagValue\\": \\"rg-aek25bwhtt22cjq\\"}]
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 2
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// example:
	//
	// low
	Severity   *string   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Stacktrace []*string `json:"Stacktrace,omitempty" xml:"Stacktrace,omitempty" type:"Repeated"`
	// example:
	//
	// 1724033134746
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 1712284140
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// exact
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// unionId。
	//
	// example:
	//
	// 23yaet4dg6ek3s7aj
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
	// example:
	//
	// https://other-general-huadong1.oss-cn-hangzhou.aliyuncs.com/uploadWidget/TaxiInvoice_01.jpeg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeAttacksResponseBodyAttacks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttacksResponseBodyAttacks) GoString() string {
	return s.String()
}

func (s *DescribeAttacksResponseBodyAttacks) GetAppDir() *string {
	return s.AppDir
}

func (s *DescribeAttacksResponseBodyAttacks) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAttacksResponseBodyAttacks) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAttacksResponseBodyAttacks) GetAvd() *string {
	return s.Avd
}

func (s *DescribeAttacksResponseBodyAttacks) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribeAttacksResponseBodyAttacks) GetConfidence() *string {
	return s.Confidence
}

func (s *DescribeAttacksResponseBodyAttacks) GetContent() *string {
	return s.Content
}

func (s *DescribeAttacksResponseBodyAttacks) GetContentLength() *int64 {
	return s.ContentLength
}

func (s *DescribeAttacksResponseBodyAttacks) GetCount() *int64 {
	return s.Count
}

func (s *DescribeAttacksResponseBodyAttacks) GetData() *string {
	return s.Data
}

func (s *DescribeAttacksResponseBodyAttacks) GetHeaders() *string {
	return s.Headers
}

func (s *DescribeAttacksResponseBodyAttacks) GetHostId() *string {
	return s.HostId
}

func (s *DescribeAttacksResponseBodyAttacks) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeAttacksResponseBodyAttacks) GetInputParamItemList() []*DescribeAttacksResponseBodyAttacksInputParamItemList {
	return s.InputParamItemList
}

func (s *DescribeAttacksResponseBodyAttacks) GetInstallType() *int64 {
	return s.InstallType
}

func (s *DescribeAttacksResponseBodyAttacks) GetIp() *string {
	return s.Ip
}

func (s *DescribeAttacksResponseBodyAttacks) GetJdk() *string {
	return s.Jdk
}

func (s *DescribeAttacksResponseBodyAttacks) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAttacksResponseBodyAttacks) GetLanguageVersion() *string {
	return s.LanguageVersion
}

func (s *DescribeAttacksResponseBodyAttacks) GetMaxTimestamp() *int64 {
	return s.MaxTimestamp
}

func (s *DescribeAttacksResponseBodyAttacks) GetMessage() *string {
	return s.Message
}

func (s *DescribeAttacksResponseBodyAttacks) GetMethod() *string {
	return s.Method
}

func (s *DescribeAttacksResponseBodyAttacks) GetMiddlewareInstanceId() *string {
	return s.MiddlewareInstanceId
}

func (s *DescribeAttacksResponseBodyAttacks) GetMinTimestamp() *int64 {
	return s.MinTimestamp
}

func (s *DescribeAttacksResponseBodyAttacks) GetOs() *string {
	return s.Os
}

func (s *DescribeAttacksResponseBodyAttacks) GetOsArch() *string {
	return s.OsArch
}

func (s *DescribeAttacksResponseBodyAttacks) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DescribeAttacksResponseBodyAttacks) GetParam() *string {
	return s.Param
}

func (s *DescribeAttacksResponseBodyAttacks) GetPayload() *string {
	return s.Payload
}

func (s *DescribeAttacksResponseBodyAttacks) GetPayloadLength() *int64 {
	return s.PayloadLength
}

func (s *DescribeAttacksResponseBodyAttacks) GetPid() *string {
	return s.Pid
}

func (s *DescribeAttacksResponseBodyAttacks) GetRaspVersion() *string {
	return s.RaspVersion
}

func (s *DescribeAttacksResponseBodyAttacks) GetRegion() *string {
	return s.Region
}

func (s *DescribeAttacksResponseBodyAttacks) GetRemote() *string {
	return s.Remote
}

func (s *DescribeAttacksResponseBodyAttacks) GetResult() *string {
	return s.Result
}

func (s *DescribeAttacksResponseBodyAttacks) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeAttacksResponseBodyAttacks) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeAttacksResponseBodyAttacks) GetStacktrace() []*string {
	return s.Stacktrace
}

func (s *DescribeAttacksResponseBodyAttacks) GetTime() *string {
	return s.Time
}

func (s *DescribeAttacksResponseBodyAttacks) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeAttacksResponseBodyAttacks) GetType() *string {
	return s.Type
}

func (s *DescribeAttacksResponseBodyAttacks) GetUnionId() *string {
	return s.UnionId
}

func (s *DescribeAttacksResponseBodyAttacks) GetUrl() *string {
	return s.Url
}

func (s *DescribeAttacksResponseBodyAttacks) SetAppDir(v string) *DescribeAttacksResponseBodyAttacks {
	s.AppDir = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetAppId(v string) *DescribeAttacksResponseBodyAttacks {
	s.AppId = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetAppName(v string) *DescribeAttacksResponseBodyAttacks {
	s.AppName = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetAvd(v string) *DescribeAttacksResponseBodyAttacks {
	s.Avd = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetCmdline(v string) *DescribeAttacksResponseBodyAttacks {
	s.Cmdline = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetConfidence(v string) *DescribeAttacksResponseBodyAttacks {
	s.Confidence = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetContent(v string) *DescribeAttacksResponseBodyAttacks {
	s.Content = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetContentLength(v int64) *DescribeAttacksResponseBodyAttacks {
	s.ContentLength = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetCount(v int64) *DescribeAttacksResponseBodyAttacks {
	s.Count = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetData(v string) *DescribeAttacksResponseBodyAttacks {
	s.Data = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetHeaders(v string) *DescribeAttacksResponseBodyAttacks {
	s.Headers = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetHostId(v string) *DescribeAttacksResponseBodyAttacks {
	s.HostId = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetHostname(v string) *DescribeAttacksResponseBodyAttacks {
	s.Hostname = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetInputParamItemList(v []*DescribeAttacksResponseBodyAttacksInputParamItemList) *DescribeAttacksResponseBodyAttacks {
	s.InputParamItemList = v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetInstallType(v int64) *DescribeAttacksResponseBodyAttacks {
	s.InstallType = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetIp(v string) *DescribeAttacksResponseBodyAttacks {
	s.Ip = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetJdk(v string) *DescribeAttacksResponseBodyAttacks {
	s.Jdk = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetLanguage(v string) *DescribeAttacksResponseBodyAttacks {
	s.Language = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetLanguageVersion(v string) *DescribeAttacksResponseBodyAttacks {
	s.LanguageVersion = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetMaxTimestamp(v int64) *DescribeAttacksResponseBodyAttacks {
	s.MaxTimestamp = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetMessage(v string) *DescribeAttacksResponseBodyAttacks {
	s.Message = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetMethod(v string) *DescribeAttacksResponseBodyAttacks {
	s.Method = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetMiddlewareInstanceId(v string) *DescribeAttacksResponseBodyAttacks {
	s.MiddlewareInstanceId = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetMinTimestamp(v int64) *DescribeAttacksResponseBodyAttacks {
	s.MinTimestamp = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetOs(v string) *DescribeAttacksResponseBodyAttacks {
	s.Os = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetOsArch(v string) *DescribeAttacksResponseBodyAttacks {
	s.OsArch = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetOsVersion(v string) *DescribeAttacksResponseBodyAttacks {
	s.OsVersion = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetParam(v string) *DescribeAttacksResponseBodyAttacks {
	s.Param = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetPayload(v string) *DescribeAttacksResponseBodyAttacks {
	s.Payload = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetPayloadLength(v int64) *DescribeAttacksResponseBodyAttacks {
	s.PayloadLength = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetPid(v string) *DescribeAttacksResponseBodyAttacks {
	s.Pid = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetRaspVersion(v string) *DescribeAttacksResponseBodyAttacks {
	s.RaspVersion = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetRegion(v string) *DescribeAttacksResponseBodyAttacks {
	s.Region = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetRemote(v string) *DescribeAttacksResponseBodyAttacks {
	s.Remote = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetResult(v string) *DescribeAttacksResponseBodyAttacks {
	s.Result = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetRuleResult(v string) *DescribeAttacksResponseBodyAttacks {
	s.RuleResult = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetSeverity(v string) *DescribeAttacksResponseBodyAttacks {
	s.Severity = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetStacktrace(v []*string) *DescribeAttacksResponseBodyAttacks {
	s.Stacktrace = v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetTime(v string) *DescribeAttacksResponseBodyAttacks {
	s.Time = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetTimestamp(v int64) *DescribeAttacksResponseBodyAttacks {
	s.Timestamp = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetType(v string) *DescribeAttacksResponseBodyAttacks {
	s.Type = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetUnionId(v string) *DescribeAttacksResponseBodyAttacks {
	s.UnionId = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) SetUrl(v string) *DescribeAttacksResponseBodyAttacks {
	s.Url = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacks) Validate() error {
	return dara.Validate(s)
}

type DescribeAttacksResponseBodyAttacksInputParamItemList struct {
	// example:
	//
	// 4E7772EA-1AD0-5583-BA55-A4****
	ProcessedKey *string `json:"ProcessedKey,omitempty" xml:"ProcessedKey,omitempty"`
	// example:
	//
	// 4E7772EA-1AD0-5583-BA55-A4****
	RawKey *string `json:"RawKey,omitempty" xml:"RawKey,omitempty"`
	// example:
	//
	// 10.101.9.189
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAttacksResponseBodyAttacksInputParamItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttacksResponseBodyAttacksInputParamItemList) GoString() string {
	return s.String()
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) GetProcessedKey() *string {
	return s.ProcessedKey
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) GetRawKey() *string {
	return s.RawKey
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) GetValue() *string {
	return s.Value
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) SetProcessedKey(v string) *DescribeAttacksResponseBodyAttacksInputParamItemList {
	s.ProcessedKey = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) SetRawKey(v string) *DescribeAttacksResponseBodyAttacksInputParamItemList {
	s.RawKey = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) SetValue(v string) *DescribeAttacksResponseBodyAttacksInputParamItemList {
	s.Value = &v
	return s
}

func (s *DescribeAttacksResponseBodyAttacksInputParamItemList) Validate() error {
	return dara.Validate(s)
}
