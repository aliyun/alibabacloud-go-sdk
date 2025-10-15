// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDialogAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RunDialogAnalysisResponseBody
	GetCost() *int64
	SetData(v *RunDialogAnalysisResponseBodyData) *RunDialogAnalysisResponseBody
	GetData() *RunDialogAnalysisResponseBodyData
	SetDataType(v string) *RunDialogAnalysisResponseBody
	GetDataType() *string
	SetErrCode(v string) *RunDialogAnalysisResponseBody
	GetErrCode() *string
	SetMessage(v string) *RunDialogAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunDialogAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunDialogAnalysisResponseBody
	GetSuccess() *bool
	SetTime(v string) *RunDialogAnalysisResponseBody
	GetTime() *string
}

type RunDialogAnalysisResponseBody struct {
	// example:
	//
	// null
	Cost *int64                             `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RunDialogAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 02CD4454-3F2C-57D0-9060-68DEAA1F6993
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunDialogAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RunDialogAnalysisResponseBody) GetData() *RunDialogAnalysisResponseBodyData {
	return s.Data
}

func (s *RunDialogAnalysisResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RunDialogAnalysisResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RunDialogAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunDialogAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDialogAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunDialogAnalysisResponseBody) GetTime() *string {
	return s.Time
}

func (s *RunDialogAnalysisResponseBody) SetCost(v int64) *RunDialogAnalysisResponseBody {
	s.Cost = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetData(v *RunDialogAnalysisResponseBodyData) *RunDialogAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetDataType(v string) *RunDialogAnalysisResponseBody {
	s.DataType = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetErrCode(v string) *RunDialogAnalysisResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetMessage(v string) *RunDialogAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetRequestId(v string) *RunDialogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetSuccess(v bool) *RunDialogAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) SetTime(v string) *RunDialogAnalysisResponseBody {
	s.Time = &v
	return s
}

func (s *RunDialogAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDialogAnalysisResponseBodyData struct {
	DialogAnalysisRespList []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespList `json:"dialogAnalysisRespList,omitempty" xml:"dialogAnalysisRespList,omitempty" type:"Repeated"`
}

func (s RunDialogAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponseBodyData) GetDialogAnalysisRespList() []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	return s.DialogAnalysisRespList
}

func (s *RunDialogAnalysisResponseBodyData) SetDialogAnalysisRespList(v []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) *RunDialogAnalysisResponseBodyData {
	s.DialogAnalysisRespList = v
	return s
}

func (s *RunDialogAnalysisResponseBodyData) Validate() error {
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

type RunDialogAnalysisResponseBodyDataDialogAnalysisRespList struct {
	AnalysisResp *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp `json:"analysisResp,omitempty" xml:"analysisResp,omitempty" type:"Struct"`
	FailNode     []*string                                                            `json:"failNode,omitempty" xml:"failNode,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-04-24 11:54:34
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1759457905S001vejpvd6vej
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GetAnalysisResp() *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	return s.AnalysisResp
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GetFailNode() []*string {
	return s.FailNode
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) GetStatus() *string {
	return s.Status
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) SetAnalysisResp(v *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	s.AnalysisResp = v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) SetFailNode(v []*string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	s.FailNode = v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) SetGmtCreate(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	s.GmtCreate = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) SetSessionId(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	s.SessionId = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) SetStatus(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList {
	s.Status = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespList) Validate() error {
	if s.AnalysisResp != nil {
		if err := s.AnalysisResp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp struct {
	DialogExecPlan        *string                                                                            `json:"dialogExecPlan,omitempty" xml:"dialogExecPlan,omitempty"`
	DialogLabels          []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels `json:"dialogLabels,omitempty" xml:"dialogLabels,omitempty" type:"Repeated"`
	DialogOpenAnalysis    map[string]interface{}                                                             `json:"dialogOpenAnalysis,omitempty" xml:"dialogOpenAnalysis,omitempty"`
	DialogProcessAnalysis map[string]interface{}                                                             `json:"dialogProcessAnalysis,omitempty" xml:"dialogProcessAnalysis,omitempty"`
	DialogSop             *string                                                                            `json:"dialogSop,omitempty" xml:"dialogSop,omitempty"`
	DialogSummary         *string                                                                            `json:"dialogSummary,omitempty" xml:"dialogSummary,omitempty"`
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogExecPlan() *string {
	return s.DialogExecPlan
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogLabels() []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	return s.DialogLabels
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogOpenAnalysis() map[string]interface{} {
	return s.DialogOpenAnalysis
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogProcessAnalysis() map[string]interface{} {
	return s.DialogProcessAnalysis
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogSop() *string {
	return s.DialogSop
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) GetDialogSummary() *string {
	return s.DialogSummary
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogExecPlan(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogExecPlan = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogLabels(v []*RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogLabels = v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogOpenAnalysis(v map[string]interface{}) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogOpenAnalysis = v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogProcessAnalysis(v map[string]interface{}) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogProcessAnalysis = v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSop(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSop = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSummary(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSummary = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisResp) Validate() error {
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

type RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GetName() *string {
	return s.Name
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GetValue() *string {
	return s.Value
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetName(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Name = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetValue(v string) *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Value = &v
	return s
}

func (s *RunDialogAnalysisResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) Validate() error {
	return dara.Validate(s)
}
