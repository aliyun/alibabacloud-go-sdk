// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryRasrEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryRasrEventResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryRasrEventResponseBody
	GetCode() *string
	SetData(v *CloudQueryRasrEventResponseBodyData) *CloudQueryRasrEventResponseBody
	GetData() *CloudQueryRasrEventResponseBodyData
	SetMessage(v string) *CloudQueryRasrEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryRasrEventResponseBody
	GetRequestId() *string
}

type CloudQueryRasrEventResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryRasrEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryRasrEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryRasrEventResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryRasrEventResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryRasrEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryRasrEventResponseBody) GetData() *CloudQueryRasrEventResponseBodyData {
	return s.Data
}

func (s *CloudQueryRasrEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryRasrEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryRasrEventResponseBody) SetAccessDeniedDetail(v string) *CloudQueryRasrEventResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryRasrEventResponseBody) SetCode(v string) *CloudQueryRasrEventResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryRasrEventResponseBody) SetData(v *CloudQueryRasrEventResponseBodyData) *CloudQueryRasrEventResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryRasrEventResponseBody) SetMessage(v string) *CloudQueryRasrEventResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryRasrEventResponseBody) SetRequestId(v string) *CloudQueryRasrEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryRasrEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryRasrEventResponseBodyData struct {
	// 返回数据
	List []*CloudQueryRasrEventResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryRasrEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryRasrEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryRasrEventResponseBodyData) GetList() []*CloudQueryRasrEventResponseBodyDataList {
	return s.List
}

func (s *CloudQueryRasrEventResponseBodyData) SetList(v []*CloudQueryRasrEventResponseBodyDataList) *CloudQueryRasrEventResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryRasrEventResponseBodyData) Validate() error {
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

type CloudQueryRasrEventResponseBodyDataList struct {
	// 机器人对话文本信息
	//
	// example:
	//
	// 示例值
	BotText *string `json:"BotText,omitempty" xml:"BotText,omitempty"`
	// 坐席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 177xxxx1313
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1768615738
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 引擎
	//
	// example:
	//
	// ""
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// 呼叫中心 id
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// sip-2-1558615724.87
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 混音录音文件名
	//
	// example:
	//
	// sip-2-1558615724.87-54529-mix
	MixRecordFile *string `json:"MixRecordFile,omitempty" xml:"MixRecordFile,omitempty"`
	// 1:座席侧（webcall是第二侧），2:客户侧（webcall是第一侧）
	//
	// example:
	//
	// 2
	MonitorSide *string `json:"MonitorSide,omitempty" xml:"MonitorSide,omitempty"`
	// 机器人侧录音文件名
	//
	// example:
	//
	// sip-2-1558615724.87-54529-out
	OutRecordFile *string `json:"OutRecordFile,omitempty" xml:"OutRecordFile,omitempty"`
	// 持续时间
	//
	// example:
	//
	// 333
	RasrDuration *string `json:"RasrDuration,omitempty" xml:"RasrDuration,omitempty"`
	// 客户侧录音文件名
	//
	// example:
	//
	// sip-2-1558615724.87-54529-in
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 96
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 客户对话文本信息
	//
	// example:
	//
	// 示例值示例值示例值
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// sip-2-1480660149.0
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s CloudQueryRasrEventResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryRasrEventResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetBotText() *string {
	return s.BotText
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetEngine() *string {
	return s.Engine
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetMixRecordFile() *string {
	return s.MixRecordFile
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetMonitorSide() *string {
	return s.MonitorSide
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetOutRecordFile() *string {
	return s.OutRecordFile
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetRasrDuration() *string {
	return s.RasrDuration
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetRecordFile() *string {
	return s.RecordFile
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetText() *string {
	return s.Text
}

func (s *CloudQueryRasrEventResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetBotText(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.BotText = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetCno(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetCustomerNumber(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetEndTime(v int64) *CloudQueryRasrEventResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetEngine(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.Engine = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetEnterpriseId(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetMainUniqueId(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.MainUniqueId = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetMixRecordFile(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.MixRecordFile = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetMonitorSide(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.MonitorSide = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetOutRecordFile(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.OutRecordFile = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetRasrDuration(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.RasrDuration = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetRecordFile(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.RecordFile = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetStartTime(v int64) *CloudQueryRasrEventResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetText(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.Text = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) SetUniqueId(v string) *CloudQueryRasrEventResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryRasrEventResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
