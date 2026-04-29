// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListCurlLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CloudListCurlLogRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudListCurlLogRequest
	GetEnterpriseId() *int64
	SetLimit(v string) *CloudListCurlLogRequest
	GetLimit() *string
	SetResult(v int64) *CloudListCurlLogRequest
	GetResult() *int64
	SetRetry(v int64) *CloudListCurlLogRequest
	GetRetry() *int64
	SetStart(v string) *CloudListCurlLogRequest
	GetStart() *string
	SetStartTime(v string) *CloudListCurlLogRequest
	GetStartTime() *string
	SetType(v int64) *CloudListCurlLogRequest
	GetType() *int64
	SetUrl(v string) *CloudListCurlLogRequest
	GetUrl() *string
}

type CloudListCurlLogRequest struct {
	// 查询结束时间；格式:"yyyy-MM-dd HH:mm:ss",精确到s。startTime与endTime不允许跨月，默认值取当前月份最后一天
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-23 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 需要取出的数据条数；大于0，最大不能超过500，默认10
	//
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 推送结果；取值：1、成功，0、失败
	//
	// example:
	//
	// 1
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// 重试次数；默认-2
	//
	// example:
	//
	// 3
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 从第几条开始取；大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// 查询开始时间；格式:"yyyy-MM-dd HH:mm:ss",精确到s。startTime与endTime不允许跨月，默认值取当前月份第一天
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 推送类型；1:来电推送、2:来电响铃推送、3:外呼响铃推送、4:webcall转座席响铃推送、5:来电接通推送、6:外呼接通推送、7:来电挂机推送、8:外呼挂机推送、9:座席状态推送、10:按键推送、11:号码状态识别推送、12:录音状态推送、13:HTTP交互推送、15:ASR语音转换结果推送、14:预测式外呼任务推送、16:webcall客户侧响铃推送、17:满意度调查推送、18:预测外呼客户接听状态推送、19:NLU推送、20:预测外呼双方接听推送、21:IVR交互结束推送、22:外呼客户响铃推送、23:webcall客户接听推送、24:webcall转接响铃推送、25:webcall转接接通推送、26:主叫外呼分机响铃推送、27:主叫外呼分机接听推送、28:主叫外呼客户响铃推送、29:主叫外呼双方接通推送
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 推送的url；说明：目的URL
	//
	// example:
	//
	// ""
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CloudListCurlLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListCurlLogRequest) GoString() string {
	return s.String()
}

func (s *CloudListCurlLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudListCurlLogRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListCurlLogRequest) GetLimit() *string {
	return s.Limit
}

func (s *CloudListCurlLogRequest) GetResult() *int64 {
	return s.Result
}

func (s *CloudListCurlLogRequest) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudListCurlLogRequest) GetStart() *string {
	return s.Start
}

func (s *CloudListCurlLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudListCurlLogRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudListCurlLogRequest) GetUrl() *string {
	return s.Url
}

func (s *CloudListCurlLogRequest) SetEndTime(v string) *CloudListCurlLogRequest {
	s.EndTime = &v
	return s
}

func (s *CloudListCurlLogRequest) SetEnterpriseId(v int64) *CloudListCurlLogRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListCurlLogRequest) SetLimit(v string) *CloudListCurlLogRequest {
	s.Limit = &v
	return s
}

func (s *CloudListCurlLogRequest) SetResult(v int64) *CloudListCurlLogRequest {
	s.Result = &v
	return s
}

func (s *CloudListCurlLogRequest) SetRetry(v int64) *CloudListCurlLogRequest {
	s.Retry = &v
	return s
}

func (s *CloudListCurlLogRequest) SetStart(v string) *CloudListCurlLogRequest {
	s.Start = &v
	return s
}

func (s *CloudListCurlLogRequest) SetStartTime(v string) *CloudListCurlLogRequest {
	s.StartTime = &v
	return s
}

func (s *CloudListCurlLogRequest) SetType(v int64) *CloudListCurlLogRequest {
	s.Type = &v
	return s
}

func (s *CloudListCurlLogRequest) SetUrl(v string) *CloudListCurlLogRequest {
	s.Url = &v
	return s
}

func (s *CloudListCurlLogRequest) Validate() error {
	return dara.Validate(s)
}
