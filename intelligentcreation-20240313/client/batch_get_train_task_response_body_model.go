// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetTrainTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetTrainTaskResponseBody
	GetRequestId() *string
	SetVoiceList(v []*BatchGetTrainTaskResponseBodyVoiceList) *BatchGetTrainTaskResponseBody
	GetVoiceList() []*BatchGetTrainTaskResponseBodyVoiceList
}

type BatchGetTrainTaskResponseBody struct {
	// example:
	//
	// 2226A26A-26E5-5AB9-A14A-54D612FCF96A
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VoiceList []*BatchGetTrainTaskResponseBodyVoiceList `json:"voiceList,omitempty" xml:"voiceList,omitempty" type:"Repeated"`
}

func (s BatchGetTrainTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetTrainTaskResponseBody) GetVoiceList() []*BatchGetTrainTaskResponseBodyVoiceList {
	return s.VoiceList
}

func (s *BatchGetTrainTaskResponseBody) SetRequestId(v string) *BatchGetTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBody) SetVoiceList(v []*BatchGetTrainTaskResponseBodyVoiceList) *BatchGetTrainTaskResponseBody {
	s.VoiceList = v
	return s
}

func (s *BatchGetTrainTaskResponseBody) Validate() error {
	if s.VoiceList != nil {
		for _, item := range s.VoiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetTrainTaskResponseBodyVoiceList struct {
	// example:
	//
	// 1524004782438111
	AliyunSubId      *string `json:"aliyunSubId,omitempty" xml:"aliyunSubId,omitempty"`
	AuditFailMessage *string `json:"auditFailMessage,omitempty" xml:"auditFailMessage,omitempty"`
	// example:
	//
	// auditFail
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// M
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// BASIC_MODEL
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// CopyVoice
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TrainFailMessage *string `json:"trainFailMessage,omitempty" xml:"trainFailMessage,omitempty"`
	// example:
	//
	// trainFail
	TrainStatus *string `json:"trainStatus,omitempty" xml:"trainStatus,omitempty"`
	// example:
	//
	// realTimeInteractivity
	UseScene      *string                                              `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceMaterial *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial `json:"voiceMaterial,omitempty" xml:"voiceMaterial,omitempty" type:"Struct"`
}

func (s BatchGetTrainTaskResponseBodyVoiceList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskResponseBodyVoiceList) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetAliyunSubId() *string {
	return s.AliyunSubId
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetAuditFailMessage() *string {
	return s.AuditFailMessage
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetGender() *string {
	return s.Gender
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetName() *string {
	return s.Name
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetTaskType() *string {
	return s.TaskType
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetTrainFailMessage() *string {
	return s.TrainFailMessage
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetTrainStatus() *string {
	return s.TrainStatus
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetUseScene() *string {
	return s.UseScene
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) GetVoiceMaterial() *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	return s.VoiceMaterial
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAliyunSubId(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AliyunSubId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAuditFailMessage(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AuditFailMessage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAuditStatus(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AuditStatus = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetCreateTime(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.CreateTime = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetGender(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.Gender = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetName(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.Name = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetResSpecType(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.ResSpecType = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTaskId(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TaskId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTaskType(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TaskType = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTrainFailMessage(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TrainFailMessage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTrainStatus(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TrainStatus = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetUseScene(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.UseScene = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetVoiceMaterial(v *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) *BatchGetTrainTaskResponseBodyVoiceList {
	s.VoiceMaterial = v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) Validate() error {
	if s.VoiceMaterial != nil {
		if err := s.VoiceMaterial.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial struct {
	// example:
	//
	// 1
	VoiceId *int64 `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
	// example:
	//
	// zh
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	// example:
	//
	// http://www.voice.com
	VoiceUrl *string `json:"voiceUrl,omitempty" xml:"voiceUrl,omitempty"`
}

func (s BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) GetVoiceId() *int64 {
	return s.VoiceId
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) GetVoiceUrl() *string {
	return s.VoiceUrl
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceId(v int64) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceLanguage(v string) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceLanguage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceUrl(v string) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceUrl = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) Validate() error {
	return dara.Validate(s)
}
