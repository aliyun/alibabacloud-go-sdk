// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListCurlLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListCurlLogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListCurlLogResponseBody
	GetCode() *string
	SetData(v *CloudListCurlLogResponseBodyData) *CloudListCurlLogResponseBody
	GetData() *CloudListCurlLogResponseBodyData
	SetMessage(v string) *CloudListCurlLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListCurlLogResponseBody
	GetRequestId() *string
}

type CloudListCurlLogResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListCurlLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListCurlLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListCurlLogResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListCurlLogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListCurlLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListCurlLogResponseBody) GetData() *CloudListCurlLogResponseBodyData {
	return s.Data
}

func (s *CloudListCurlLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListCurlLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListCurlLogResponseBody) SetAccessDeniedDetail(v string) *CloudListCurlLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListCurlLogResponseBody) SetCode(v string) *CloudListCurlLogResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListCurlLogResponseBody) SetData(v *CloudListCurlLogResponseBodyData) *CloudListCurlLogResponseBody {
	s.Data = v
	return s
}

func (s *CloudListCurlLogResponseBody) SetMessage(v string) *CloudListCurlLogResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListCurlLogResponseBody) SetRequestId(v string) *CloudListCurlLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListCurlLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListCurlLogResponseBodyData struct {
	// 返回数据JSON格式
	List []*CloudListCurlLogResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListCurlLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListCurlLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListCurlLogResponseBodyData) GetList() []*CloudListCurlLogResponseBodyDataList {
	return s.List
}

func (s *CloudListCurlLogResponseBodyData) SetList(v []*CloudListCurlLogResponseBodyDataList) *CloudListCurlLogResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListCurlLogResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListCurlLogResponseBodyDataList struct {
	// 文本类型
	//
	// example:
	//
	// 1
	ContentType *int64 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 1774822134415
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 当月
	//
	// example:
	//
	// ""
	CurrentMonth *string `json:"CurrentMonth,omitempty" xml:"CurrentMonth,omitempty"`
	// 延迟推送时间
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// 推送结束时间
	//
	// example:
	//
	// 1774822137959
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 企业id
	//
	// example:
	//
	// 7122600
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 请求头
	//
	// example:
	//
	// {}
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// 日志id
	//
	// example:
	//
	// 12715
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 推送登记
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 推送方法
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// 推送参数
	//
	// example:
	//
	// 示例值示例值示例值
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 推送位置
	//
	// example:
	//
	// 0
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// 推送时间
	//
	// example:
	//
	// 1774822137191
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// 推送状态
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 返回文本
	//
	// example:
	//
	// 示例值示例值
	ResultText *string `json:"ResultText,omitempty" xml:"ResultText,omitempty"`
	// 重试次数
	//
	// example:
	//
	// 0
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 推送开始时间
	//
	// example:
	//
	// 1774822134132
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态码
	//
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 超时时间
	//
	// example:
	//
	// 3
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// 推送类型
	//
	// example:
	//
	// 9
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 推送类型字符串
	//
	// example:
	//
	// 示例值
	TypeStr *string `json:"TypeStr,omitempty" xml:"TypeStr,omitempty"`
	// 唯一标示
	//
	// example:
	//
	// ""
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// 推送url
	//
	// example:
	//
	// https://www.aaa.com/index.htm
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CloudListCurlLogResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListCurlLogResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListCurlLogResponseBodyDataList) GetContentType() *int64 {
	return s.ContentType
}

func (s *CloudListCurlLogResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListCurlLogResponseBodyDataList) GetCurrentMonth() *string {
	return s.CurrentMonth
}

func (s *CloudListCurlLogResponseBodyDataList) GetDelay() *int64 {
	return s.Delay
}

func (s *CloudListCurlLogResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudListCurlLogResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudListCurlLogResponseBodyDataList) GetHeaders() *string {
	return s.Headers
}

func (s *CloudListCurlLogResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudListCurlLogResponseBodyDataList) GetLevel() *int64 {
	return s.Level
}

func (s *CloudListCurlLogResponseBodyDataList) GetMethod() *string {
	return s.Method
}

func (s *CloudListCurlLogResponseBodyDataList) GetParams() *string {
	return s.Params
}

func (s *CloudListCurlLogResponseBodyDataList) GetPosition() *string {
	return s.Position
}

func (s *CloudListCurlLogResponseBodyDataList) GetRequestTime() *string {
	return s.RequestTime
}

func (s *CloudListCurlLogResponseBodyDataList) GetResult() *string {
	return s.Result
}

func (s *CloudListCurlLogResponseBodyDataList) GetResultText() *string {
	return s.ResultText
}

func (s *CloudListCurlLogResponseBodyDataList) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudListCurlLogResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudListCurlLogResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudListCurlLogResponseBodyDataList) GetTimeout() *string {
	return s.Timeout
}

func (s *CloudListCurlLogResponseBodyDataList) GetType() *int64 {
	return s.Type
}

func (s *CloudListCurlLogResponseBodyDataList) GetTypeStr() *string {
	return s.TypeStr
}

func (s *CloudListCurlLogResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudListCurlLogResponseBodyDataList) GetUrl() *string {
	return s.Url
}

func (s *CloudListCurlLogResponseBodyDataList) SetContentType(v int64) *CloudListCurlLogResponseBodyDataList {
	s.ContentType = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetCreateTime(v string) *CloudListCurlLogResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetCurrentMonth(v string) *CloudListCurlLogResponseBodyDataList {
	s.CurrentMonth = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetDelay(v int64) *CloudListCurlLogResponseBodyDataList {
	s.Delay = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetEndTime(v string) *CloudListCurlLogResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetEnterpriseId(v string) *CloudListCurlLogResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetHeaders(v string) *CloudListCurlLogResponseBodyDataList {
	s.Headers = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetId(v string) *CloudListCurlLogResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetLevel(v int64) *CloudListCurlLogResponseBodyDataList {
	s.Level = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetMethod(v string) *CloudListCurlLogResponseBodyDataList {
	s.Method = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetParams(v string) *CloudListCurlLogResponseBodyDataList {
	s.Params = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetPosition(v string) *CloudListCurlLogResponseBodyDataList {
	s.Position = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetRequestTime(v string) *CloudListCurlLogResponseBodyDataList {
	s.RequestTime = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetResult(v string) *CloudListCurlLogResponseBodyDataList {
	s.Result = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetResultText(v string) *CloudListCurlLogResponseBodyDataList {
	s.ResultText = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetRetry(v int64) *CloudListCurlLogResponseBodyDataList {
	s.Retry = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetStartTime(v string) *CloudListCurlLogResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetStatusCode(v string) *CloudListCurlLogResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetTimeout(v string) *CloudListCurlLogResponseBodyDataList {
	s.Timeout = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetType(v int64) *CloudListCurlLogResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetTypeStr(v string) *CloudListCurlLogResponseBodyDataList {
	s.TypeStr = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetUniqueId(v string) *CloudListCurlLogResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) SetUrl(v string) *CloudListCurlLogResponseBodyDataList {
	s.Url = &v
	return s
}

func (s *CloudListCurlLogResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
