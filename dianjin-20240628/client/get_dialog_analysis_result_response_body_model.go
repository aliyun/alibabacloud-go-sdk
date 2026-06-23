// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogAnalysisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDialogAnalysisResultResponseBody
	GetCost() *int64
	SetData(v *GetDialogAnalysisResultResponseBodyData) *GetDialogAnalysisResultResponseBody
	GetData() *GetDialogAnalysisResultResponseBodyData
	SetDataType(v string) *GetDialogAnalysisResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDialogAnalysisResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDialogAnalysisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDialogAnalysisResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDialogAnalysisResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDialogAnalysisResultResponseBody
	GetTime() *string
}

type GetDialogAnalysisResultResponseBody struct {
	// Processing time in milliseconds
	//
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// Response data
	Data *GetDialogAnalysisResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Data type
	//
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// Error code
	//
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// Error message
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 88A006F0-B565-53BA-B38A-DBDF9D0B2935
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// Timestamp
	//
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDialogAnalysisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDialogAnalysisResultResponseBody) GetData() *GetDialogAnalysisResultResponseBodyData {
	return s.Data
}

func (s *GetDialogAnalysisResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDialogAnalysisResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDialogAnalysisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDialogAnalysisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDialogAnalysisResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDialogAnalysisResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDialogAnalysisResultResponseBody) SetCost(v int64) *GetDialogAnalysisResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetData(v *GetDialogAnalysisResultResponseBodyData) *GetDialogAnalysisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetDataType(v string) *GetDialogAnalysisResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetErrCode(v string) *GetDialogAnalysisResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetMessage(v string) *GetDialogAnalysisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetRequestId(v string) *GetDialogAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetSuccess(v bool) *GetDialogAnalysisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetTime(v string) *GetDialogAnalysisResultResponseBody {
	s.Time = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDialogAnalysisResultResponseBodyData struct {
	// List of session analysis results
	DialogAnalysisRespList []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList `json:"dialogAnalysisRespList,omitempty" xml:"dialogAnalysisRespList,omitempty" type:"Repeated"`
}

func (s GetDialogAnalysisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyData) GetDialogAnalysisRespList() []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	return s.DialogAnalysisRespList
}

func (s *GetDialogAnalysisResultResponseBodyData) SetDialogAnalysisRespList(v []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) *GetDialogAnalysisResultResponseBodyData {
	s.DialogAnalysisRespList = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyData) Validate() error {
	if s.DialogAnalysisRespList != nil {
		for _, item := range s.DialogAnalysisRespList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList struct {
	// Session analysis result
	AnalysisResp *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp `json:"analysisResp,omitempty" xml:"analysisResp,omitempty" type:"Struct"`
	// Session creation time
	//
	// example:
	//
	// 2024-04-24 11:54:34
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// OSS URL for the session analysis result. The URL expires in one hour.
	//
	// example:
	//
	// https://xxx.oss-cn-beijing.aliyuncs.com/dialog-analysis/2024-12-30/2/1826661605606129665
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
	// Session ID
	//
	// example:
	//
	// 183764873624
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Task execution status for session analysis.
	//
	// - init means the task has not started
	//
	// - pending means the task is queued
	//
	// - running means the task is in progress
	//
	// - error means the task failed
	//
	// - success means the task completed successfully
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GetAnalysisResp() *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	return s.AnalysisResp
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GetSessionId() *string {
	return s.SessionId
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GetStatus() *string {
	return s.Status
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetAnalysisResp(v *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.AnalysisResp = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetGmtCreate(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.GmtCreate = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetOssUrl(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.OssUrl = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetSessionId(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.SessionId = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetStatus(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.Status = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) Validate() error {
	if s.AnalysisResp != nil {
		if err := s.AnalysisResp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp struct {
	// Session execution plan
	//
	// example:
	//
	// 1. 客服应再次确认客户的疑问是否已解决，特别是关于额度的具体数额。\\n2. 如果客户仍有疑问，提供客服热线电话，建议客户直接拨打以获取更详细的帮助。\\n3. 提醒客户检查短信中的链接，以便快速查看和操作。\\n4. 记录此次通话中客户表现出的任何不适或不便，确保后续跟进时更加体贴。\\n5. 发送一条包含操作指南的短信，确保客户能够轻松找到并使用服务。\\n6. 结束通话前，再次感谢客户的支持，并表达希望客户早日康复的愿望。
	DialogExecPlan *string `json:"dialogExecPlan,omitempty" xml:"dialogExecPlan,omitempty"`
	// List of session labels
	DialogLabels []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels `json:"dialogLabels,omitempty" xml:"dialogLabels,omitempty" type:"Repeated"`
	// Session open analysis
	//
	// example:
	//
	// {
	//
	//     "dialogues": [
	//
	//         {
	//
	//             "round": 1,
	//
	//             "result": [
	//
	//                 {
	//
	//                     "key": "对话主题",
	//
	//                     "value": "XX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户反应",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户反应分析",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客服话术",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "本轮客服话术修改建议",
	//
	//                     "value": "XXX"
	//
	//                 }
	//
	//             ]
	//
	//         },
	//
	//         {
	//
	//             "round": 2,
	//
	//             "result": [
	//
	//                 {
	//
	//                     "key": "对话主题",
	//
	//                     "value": "XX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户反应",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户反应分析",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客服话术",
	//
	//                     "value": "XXX"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "本轮客服话术修改建议",
	//
	//                     "value": "XXX"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     ],
	//
	//   "dialogOpenAnalysisStr":"第一轮对话：对话主题-xx##客户反应-xx##客户反应分析-xx##客服话术-xx##本轮客服话术修改建议-xx
	//
	// 第二轮对话：对话主题-xx##客户反应-xx##客户反应分析-xx##客服话术-xx##本轮客服话术修改建议-xx"
	//
	// }
	DialogOpenAnalysis map[string]interface{} `json:"dialogOpenAnalysis,omitempty" xml:"dialogOpenAnalysis,omitempty"`
	// Session process analysis
	//
	// example:
	//
	// {
	//
	//     "dialogues": [
	//
	//         {
	//
	//             "round": 1,
	//
	//             "result": [
	//
	//                 {
	//
	//                     "key": "客服",
	//
	//                     "value": "客服回应标签"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户",
	//
	//                     "value": "客户回应态度标签"
	//
	//                 }
	//
	//             ]
	//
	//         },
	//
	//         {
	//
	//             "round": 2,
	//
	//             "result": [
	//
	//                 {
	//
	//                     "key": "客服",
	//
	//                     "value": "客服回应标签"
	//
	//                 },
	//
	//                 {
	//
	//                     "key": "客户",
	//
	//                     "value": "客户回应态度标签"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     ],
	//
	// "dialogProcessAnalysisStr":"第一轮对话：客服-客服回应标签，客户-客户回应态度标签
	//
	// 第二轮对话：客服-客服回应标签，客户-客户回应态度标签"
	//
	// }
	DialogProcessAnalysis map[string]interface{} `json:"dialogProcessAnalysis,omitempty" xml:"dialogProcessAnalysis,omitempty"`
	// Session SOP
	//
	// example:
	//
	// 产品介绍
	DialogSop *string `json:"dialogSop,omitempty" xml:"dialogSop,omitempty"`
	// Session summary
	//
	// example:
	//
	// - 是否有资金需求：不确定，客户未明确表示有无资金需求。\\n- 是否有意向：不确定，客户未明确表达意向。\\n- 是否可营销：不可营销，客户对客服的多次询问未表现出兴趣，且对话中提到因不适希望减少联系。\\n- 待满足需求：客户希望了解具体的预审额度信息。
	DialogSummary *string `json:"dialogSummary,omitempty" xml:"dialogSummary,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogExecPlan() *string {
	return s.DialogExecPlan
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogLabels() []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	return s.DialogLabels
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogOpenAnalysis() map[string]interface{} {
	return s.DialogOpenAnalysis
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogProcessAnalysis() map[string]interface{} {
	return s.DialogProcessAnalysis
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogSop() *string {
	return s.DialogSop
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogSummary() *string {
	return s.DialogSummary
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogExecPlan(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogExecPlan = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogLabels(v []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogLabels = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogOpenAnalysis(v map[string]interface{}) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogOpenAnalysis = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogProcessAnalysis(v map[string]interface{}) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogProcessAnalysis = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSop(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSop = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSummary(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSummary = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) Validate() error {
	if s.DialogLabels != nil {
		for _, item := range s.DialogLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels struct {
	// Label name
	//
	// example:
	//
	// 额度不足
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Label value
	//
	// example:
	//
	// 0
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GetName() *string {
	return s.Name
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GetValue() *string {
	return s.Value
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetName(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Name = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetValue(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Value = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) Validate() error {
	return dara.Validate(s)
}
