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
	// example:
	//
	// 9b57f0fcf98181df8d8487d1cc91cb8d
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// example:
	//
	// fc312aa0c32ba8a6147db6221fb1c1ee
	UniqueInfo    *string                                             `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	UniqueTagList []*GetAlertRecordAnalysisResultRequestUniqueTagList `json:"UniqueTagList,omitempty" xml:"UniqueTagList,omitempty" type:"Repeated"`
	// example:
	//
	// ebde6d4e3e4aba728962eec43a69196e9J7tt7H47Pc
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
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
	// example:
	//
	// 10a19b654e73ff079ede61ce3f4465e0
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// example:
	//
	// false
	ChooseLike *bool `json:"ChooseLike,omitempty" xml:"ChooseLike,omitempty"`
	// example:
	//
	// pc-bp19up785757dz800
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// i-rj9c7d4bli38***tuym
	MachineInstanceId *string `json:"MachineInstanceId,omitempty" xml:"MachineInstanceId,omitempty"`
	// example:
	//
	// 2025-06-27 00:00:00
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// BusinessLicense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// fc312aa0c32ba8a6147db6221fb1c1ee
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// example:
	//
	// 3309e55fcb1ed8d4bc6af098e62e0353RNabnQSO1bx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
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
