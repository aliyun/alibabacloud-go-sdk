// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRecordAnalysisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAlertRecordAnalysisResultResponseBody
	GetCode() *string
	SetData(v *GetAlertRecordAnalysisResultResponseBodyData) *GetAlertRecordAnalysisResultResponseBody
	GetData() *GetAlertRecordAnalysisResultResponseBodyData
	SetMessage(v string) *GetAlertRecordAnalysisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlertRecordAnalysisResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAlertRecordAnalysisResultResponseBody
	GetSuccess() *bool
}

type GetAlertRecordAnalysisResultResponseBody struct {
	// example:
	//
	// 200
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAlertRecordAnalysisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 99D93ED4-D462-5FC5-8518-9BC1C49C7B6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlertRecordAnalysisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlertRecordAnalysisResultResponseBody) GetData() *GetAlertRecordAnalysisResultResponseBodyData {
	return s.Data
}

func (s *GetAlertRecordAnalysisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlertRecordAnalysisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertRecordAnalysisResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAlertRecordAnalysisResultResponseBody) SetCode(v string) *GetAlertRecordAnalysisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBody) SetData(v *GetAlertRecordAnalysisResultResponseBodyData) *GetAlertRecordAnalysisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBody) SetMessage(v string) *GetAlertRecordAnalysisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBody) SetRequestId(v string) *GetAlertRecordAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBody) SetSuccess(v bool) *GetAlertRecordAnalysisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRecordAnalysisResultResponseBodyData struct {
	// example:
	//
	// -
	AnalysisCode  *string                                                      `json:"AnalysisCode,omitempty" xml:"AnalysisCode,omitempty"`
	UniqueTagList []*GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList `json:"UniqueTagList,omitempty" xml:"UniqueTagList,omitempty" type:"Repeated"`
}

func (s GetAlertRecordAnalysisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultResponseBodyData) GetAnalysisCode() *string {
	return s.AnalysisCode
}

func (s *GetAlertRecordAnalysisResultResponseBodyData) GetUniqueTagList() []*GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	return s.UniqueTagList
}

func (s *GetAlertRecordAnalysisResultResponseBodyData) SetAnalysisCode(v string) *GetAlertRecordAnalysisResultResponseBodyData {
	s.AnalysisCode = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyData) SetUniqueTagList(v []*GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) *GetAlertRecordAnalysisResultResponseBodyData {
	s.UniqueTagList = v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyData) Validate() error {
	if s.UniqueTagList != nil {
		for _, item := range s.UniqueTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList struct {
	// example:
	//
	// 179deb12f25baac9b1e2909c419bcb1f
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// example:
	//
	// 1248751055158884
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// test_code
	AnalysisCode *string `json:"AnalysisCode,omitempty" xml:"AnalysisCode,omitempty"`
	// example:
	//
	// exception_alert
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// example:
	//
	// true
	ChooseLike *bool `json:"ChooseLike,omitempty" xml:"ChooseLike,omitempty"`
	// example:
	//
	// 110.22.*8.111
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// i-adadasd-a**
	MachineInstanceId *string `json:"MachineInstanceId,omitempty" xml:"MachineInstanceId,omitempty"`
	// example:
	//
	// auto_breaking
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 390317ce81d28bbbd83c05a90b39cd6c
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// example:
	//
	// bb3051ca-c0dd-4da2-91be-ea5c80926132
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetAliUid() *string {
	return s.AliUid
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetAnalysisCode() *string {
	return s.AnalysisCode
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetAnalysisResult() *string {
	return s.AnalysisResult
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetChooseLike() *bool {
	return s.ChooseLike
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetIp() *string {
	return s.Ip
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetMachineInstanceId() *string {
	return s.MachineInstanceId
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetType() *string {
	return s.Type
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) GetUuid() *string {
	return s.Uuid
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetAliUid(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.AliUid = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetAnalysisCode(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.AnalysisCode = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetAnalysisResult(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.AnalysisResult = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetChooseLike(v bool) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.ChooseLike = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetIp(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.Ip = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetMachineInstanceId(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.MachineInstanceId = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetType(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.Type = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetUniqueInfo(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.UniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) SetUuid(v string) *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList {
	s.Uuid = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponseBodyDataUniqueTagList) Validate() error {
	return dara.Validate(s)
}
