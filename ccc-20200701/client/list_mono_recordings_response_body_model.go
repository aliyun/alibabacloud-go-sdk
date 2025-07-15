// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonoRecordingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMonoRecordingsResponseBody
	GetCode() *string
	SetData(v []*ListMonoRecordingsResponseBodyData) *ListMonoRecordingsResponseBody
	GetData() []*ListMonoRecordingsResponseBodyData
	SetHttpStatusCode(v int32) *ListMonoRecordingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMonoRecordingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMonoRecordingsResponseBody
	GetRequestId() *string
}

type ListMonoRecordingsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListMonoRecordingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E836EF2A-EDAE-5310-8921-F983C967D42D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMonoRecordingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMonoRecordingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMonoRecordingsResponseBody) GetData() []*ListMonoRecordingsResponseBodyData {
	return s.Data
}

func (s *ListMonoRecordingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMonoRecordingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMonoRecordingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMonoRecordingsResponseBody) SetCode(v string) *ListMonoRecordingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetData(v []*ListMonoRecordingsResponseBodyData) *ListMonoRecordingsResponseBody {
	s.Data = v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetHttpStatusCode(v int32) *ListMonoRecordingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetMessage(v string) *ListMonoRecordingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetRequestId(v string) *ListMonoRecordingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMonoRecordingsResponseBodyData struct {
	// example:
	//
	// agent@ccc-test
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// job-25920271311543****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 5903871
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// job-25920271311543****-798f1e90-1f82-42da-914c-46580c8f4c85-1656926518491.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://ccc-v2-shanghai.oss-cn-shanghai.aliyuncs.com/ccc-test/job-25920271311543****-798f1e90-1f82-42da-914c-46580c8f4c85-1656926518491.wav?Expires=1657014031&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 26972543893791****
	RamId *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 1620259200000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMonoRecordingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMonoRecordingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMonoRecordingsResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *ListMonoRecordingsResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *ListMonoRecordingsResponseBodyData) GetDuration() *string {
	return s.Duration
}

func (s *ListMonoRecordingsResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ListMonoRecordingsResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListMonoRecordingsResponseBodyData) GetRamId() *string {
	return s.RamId
}

func (s *ListMonoRecordingsResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListMonoRecordingsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMonoRecordingsResponseBodyData) SetAgentId(v string) *ListMonoRecordingsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetAgentName(v string) *ListMonoRecordingsResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetContactId(v string) *ListMonoRecordingsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetDuration(v string) *ListMonoRecordingsResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetFileName(v string) *ListMonoRecordingsResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetFileUrl(v string) *ListMonoRecordingsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetRamId(v string) *ListMonoRecordingsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetSkillGroupId(v string) *ListMonoRecordingsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetStartTime(v string) *ListMonoRecordingsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
