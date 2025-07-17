// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallChainInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalInfo(v string) *CallChainInfo
	GetAdditionalInfo() *string
	SetAppName(v string) *CallChainInfo
	GetAppName() *string
	SetAppType(v string) *CallChainInfo
	GetAppType() *string
	SetChildren(v []*CallChainInfo) *CallChainInfo
	GetChildren() []*CallChainInfo
	SetHaveSpan(v bool) *CallChainInfo
	GetHaveSpan() *bool
	SetLogMap(v map[string]map[string]interface{}) *CallChainInfo
	GetLogMap() map[string]map[string]interface{}
	SetLogTime(v int64) *CallChainInfo
	GetLogTime() *int64
	SetParentSpanId(v string) *CallChainInfo
	GetParentSpanId() *string
	SetPid(v string) *CallChainInfo
	GetPid() *string
	SetRegionId(v string) *CallChainInfo
	GetRegionId() *string
	SetResultCode(v string) *CallChainInfo
	GetResultCode() *string
	SetRpc(v string) *CallChainInfo
	GetRpc() *string
	SetRpcId(v string) *CallChainInfo
	GetRpcId() *string
	SetRpcType(v int64) *CallChainInfo
	GetRpcType() *int64
	SetServerIp(v string) *CallChainInfo
	GetServerIp() *string
	SetSpan(v int64) *CallChainInfo
	GetSpan() *int64
	SetSpanId(v string) *CallChainInfo
	GetSpanId() *string
	SetTagMap(v map[string]*string) *CallChainInfo
	GetTagMap() map[string]*string
	SetTraceId(v string) *CallChainInfo
	GetTraceId() *string
}

type CallChainInfo struct {
	AdditionalInfo *string                           `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	AppName        *string                           `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType        *string                           `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Children       []*CallChainInfo                  `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	HaveSpan       *bool                             `json:"HaveSpan,omitempty" xml:"HaveSpan,omitempty"`
	LogMap         map[string]map[string]interface{} `json:"LogMap,omitempty" xml:"LogMap,omitempty"`
	LogTime        *int64                            `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	ParentSpanId   *string                           `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	Pid            *string                           `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId       *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResultCode     *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Rpc            *string                           `json:"Rpc,omitempty" xml:"Rpc,omitempty"`
	RpcId          *string                           `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType        *int64                            `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServerIp       *string                           `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Span           *int64                            `json:"Span,omitempty" xml:"Span,omitempty"`
	SpanId         *string                           `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagMap         map[string]*string                `json:"TagMap,omitempty" xml:"TagMap,omitempty"`
	TraceId        *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CallChainInfo) String() string {
	return dara.Prettify(s)
}

func (s CallChainInfo) GoString() string {
	return s.String()
}

func (s *CallChainInfo) GetAdditionalInfo() *string {
	return s.AdditionalInfo
}

func (s *CallChainInfo) GetAppName() *string {
	return s.AppName
}

func (s *CallChainInfo) GetAppType() *string {
	return s.AppType
}

func (s *CallChainInfo) GetChildren() []*CallChainInfo {
	return s.Children
}

func (s *CallChainInfo) GetHaveSpan() *bool {
	return s.HaveSpan
}

func (s *CallChainInfo) GetLogMap() map[string]map[string]interface{} {
	return s.LogMap
}

func (s *CallChainInfo) GetLogTime() *int64 {
	return s.LogTime
}

func (s *CallChainInfo) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *CallChainInfo) GetPid() *string {
	return s.Pid
}

func (s *CallChainInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *CallChainInfo) GetResultCode() *string {
	return s.ResultCode
}

func (s *CallChainInfo) GetRpc() *string {
	return s.Rpc
}

func (s *CallChainInfo) GetRpcId() *string {
	return s.RpcId
}

func (s *CallChainInfo) GetRpcType() *int64 {
	return s.RpcType
}

func (s *CallChainInfo) GetServerIp() *string {
	return s.ServerIp
}

func (s *CallChainInfo) GetSpan() *int64 {
	return s.Span
}

func (s *CallChainInfo) GetSpanId() *string {
	return s.SpanId
}

func (s *CallChainInfo) GetTagMap() map[string]*string {
	return s.TagMap
}

func (s *CallChainInfo) GetTraceId() *string {
	return s.TraceId
}

func (s *CallChainInfo) SetAdditionalInfo(v string) *CallChainInfo {
	s.AdditionalInfo = &v
	return s
}

func (s *CallChainInfo) SetAppName(v string) *CallChainInfo {
	s.AppName = &v
	return s
}

func (s *CallChainInfo) SetAppType(v string) *CallChainInfo {
	s.AppType = &v
	return s
}

func (s *CallChainInfo) SetChildren(v []*CallChainInfo) *CallChainInfo {
	s.Children = v
	return s
}

func (s *CallChainInfo) SetHaveSpan(v bool) *CallChainInfo {
	s.HaveSpan = &v
	return s
}

func (s *CallChainInfo) SetLogMap(v map[string]map[string]interface{}) *CallChainInfo {
	s.LogMap = v
	return s
}

func (s *CallChainInfo) SetLogTime(v int64) *CallChainInfo {
	s.LogTime = &v
	return s
}

func (s *CallChainInfo) SetParentSpanId(v string) *CallChainInfo {
	s.ParentSpanId = &v
	return s
}

func (s *CallChainInfo) SetPid(v string) *CallChainInfo {
	s.Pid = &v
	return s
}

func (s *CallChainInfo) SetRegionId(v string) *CallChainInfo {
	s.RegionId = &v
	return s
}

func (s *CallChainInfo) SetResultCode(v string) *CallChainInfo {
	s.ResultCode = &v
	return s
}

func (s *CallChainInfo) SetRpc(v string) *CallChainInfo {
	s.Rpc = &v
	return s
}

func (s *CallChainInfo) SetRpcId(v string) *CallChainInfo {
	s.RpcId = &v
	return s
}

func (s *CallChainInfo) SetRpcType(v int64) *CallChainInfo {
	s.RpcType = &v
	return s
}

func (s *CallChainInfo) SetServerIp(v string) *CallChainInfo {
	s.ServerIp = &v
	return s
}

func (s *CallChainInfo) SetSpan(v int64) *CallChainInfo {
	s.Span = &v
	return s
}

func (s *CallChainInfo) SetSpanId(v string) *CallChainInfo {
	s.SpanId = &v
	return s
}

func (s *CallChainInfo) SetTagMap(v map[string]*string) *CallChainInfo {
	s.TagMap = v
	return s
}

func (s *CallChainInfo) SetTraceId(v string) *CallChainInfo {
	s.TraceId = &v
	return s
}

func (s *CallChainInfo) Validate() error {
	return dara.Validate(s)
}
