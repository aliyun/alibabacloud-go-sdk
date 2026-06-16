// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRecordAnalysisResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultShrinkRequest
	GetAlarmUniqueInfo() *string
	SetAliyunLang(v string) *GetAlertRecordAnalysisResultShrinkRequest
	GetAliyunLang() *string
	SetUniqueInfo(v string) *GetAlertRecordAnalysisResultShrinkRequest
	GetUniqueInfo() *string
	SetUniqueTagListShrink(v string) *GetAlertRecordAnalysisResultShrinkRequest
	GetUniqueTagListShrink() *string
	SetUuid(v string) *GetAlertRecordAnalysisResultShrinkRequest
	GetUuid() *string
}

type GetAlertRecordAnalysisResultShrinkRequest struct {
	AlarmUniqueInfo     *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	AliyunLang          *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	UniqueInfo          *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	UniqueTagListShrink *string `json:"UniqueTagList,omitempty" xml:"UniqueTagList,omitempty"`
	Uuid                *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetAlertRecordAnalysisResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) GetUniqueTagListShrink() *string {
	return s.UniqueTagListShrink
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) SetAlarmUniqueInfo(v string) *GetAlertRecordAnalysisResultShrinkRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) SetAliyunLang(v string) *GetAlertRecordAnalysisResultShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) SetUniqueInfo(v string) *GetAlertRecordAnalysisResultShrinkRequest {
	s.UniqueInfo = &v
	return s
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) SetUniqueTagListShrink(v string) *GetAlertRecordAnalysisResultShrinkRequest {
	s.UniqueTagListShrink = &v
	return s
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) SetUuid(v string) *GetAlertRecordAnalysisResultShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *GetAlertRecordAnalysisResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
