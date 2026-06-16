// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRecordAnalysisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultRequest
	GetAlarmUniqueInfo() *string
	SetAliyunLang(v string) *GetAlertRecordAnalysisResultRequest
	GetAliyunLang() *string
	SetUniqueInfo(v string) *GetAlertRecordAnalysisResultRequest
	GetUniqueInfo() *string
	SetUniqueTagList(v []*GetAlertRecordAnalysisResultRequestUniqueTagList) *GetAlertRecordAnalysisResultRequest
	GetUniqueTagList() []*GetAlertRecordAnalysisResultRequestUniqueTagList
	SetUuid(v string) *GetAlertRecordAnalysisResultRequest
	GetUuid() *string
}

type GetAlertRecordAnalysisResultRequest struct {
	AlarmUniqueInfo *string                                             `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	AliyunLang      *string                                             `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	UniqueInfo      *string                                             `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	UniqueTagList   []*GetAlertRecordAnalysisResultRequestUniqueTagList `json:"UniqueTagList,omitempty" xml:"UniqueTagList,omitempty" type:"Repeated"`
	Uuid            *string                                             `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAlertRecordAnalysisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *GetAlertRecordAnalysisResultRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetAlertRecordAnalysisResultRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *GetAlertRecordAnalysisResultRequest) GetUniqueTagList() []*GetAlertRecordAnalysisResultRequestUniqueTagList {
	return s.UniqueTagList
}

func (s *GetAlertRecordAnalysisResultRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAlertRecordAnalysisResultRequest) SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequest) SetAliyunLang(v string) *GetAlertRecordAnalysisResultRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequest) SetUniqueInfo(v string) *GetAlertRecordAnalysisResultRequest {
	s.UniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequest) SetUniqueTagList(v []*GetAlertRecordAnalysisResultRequestUniqueTagList) *GetAlertRecordAnalysisResultRequest {
	s.UniqueTagList = v
	return s
}

func (s *GetAlertRecordAnalysisResultRequest) SetUuid(v string) *GetAlertRecordAnalysisResultRequest {
	s.Uuid = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequest) Validate() error {
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

type GetAlertRecordAnalysisResultRequestUniqueTagList struct {
	AlarmUniqueInfo   *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	ChooseLike        *bool   `json:"ChooseLike,omitempty" xml:"ChooseLike,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MachineInstanceId *string `json:"MachineInstanceId,omitempty" xml:"MachineInstanceId,omitempty"`
	QueryTime         *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UniqueInfo        *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAlertRecordAnalysisResultRequestUniqueTagList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultRequestUniqueTagList) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetChooseLike() *bool {
	return s.ChooseLike
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetIp() *string {
	return s.Ip
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetMachineInstanceId() *string {
	return s.MachineInstanceId
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetQueryTime() *string {
	return s.QueryTime
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetType() *string {
	return s.Type
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) GetUuid() *string {
	return s.Uuid
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetChooseLike(v bool) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.ChooseLike = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetIp(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.Ip = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetMachineInstanceId(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.MachineInstanceId = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetQueryTime(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.QueryTime = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetType(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.Type = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetUniqueInfo(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.UniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) SetUuid(v string) *GetAlertRecordAnalysisResultRequestUniqueTagList {
	s.Uuid = &v
	return s
}

func (s *GetAlertRecordAnalysisResultRequestUniqueTagList) Validate() error {
	return dara.Validate(s)
}
