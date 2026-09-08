// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiChannelRecordingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMultiChannelRecordingsResponseBody
	GetCode() *string
	SetData(v []*ListMultiChannelRecordingsResponseBodyData) *ListMultiChannelRecordingsResponseBody
	GetData() []*ListMultiChannelRecordingsResponseBodyData
	SetHttpStatusCode(v int32) *ListMultiChannelRecordingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMultiChannelRecordingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMultiChannelRecordingsResponseBody
	GetRequestId() *string
}

type ListMultiChannelRecordingsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Recording list.
	Data []*ListMultiChannelRecordingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B19CD719-9F65-56A6-8B79-DA4282EA4797
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMultiChannelRecordingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiChannelRecordingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiChannelRecordingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMultiChannelRecordingsResponseBody) GetData() []*ListMultiChannelRecordingsResponseBodyData {
	return s.Data
}

func (s *ListMultiChannelRecordingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMultiChannelRecordingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMultiChannelRecordingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiChannelRecordingsResponseBody) SetCode(v string) *ListMultiChannelRecordingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBody) SetData(v []*ListMultiChannelRecordingsResponseBodyData) *ListMultiChannelRecordingsResponseBody {
	s.Data = v
	return s
}

func (s *ListMultiChannelRecordingsResponseBody) SetHttpStatusCode(v int32) *ListMultiChannelRecordingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBody) SetMessage(v string) *ListMultiChannelRecordingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBody) SetRequestId(v string) *ListMultiChannelRecordingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiChannelRecordingsResponseBodyData struct {
	// Agent call channel ID.
	//
	// example:
	//
	// ch-user-8526899****-8602****-1656926504363-job-25920271311543****
	AgentChannelId *string `json:"AgentChannelId,omitempty" xml:"AgentChannelId,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Agent name.
	//
	// example:
	//
	// 坐席小王
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Call ID.
	//
	// example:
	//
	// job-25920271311543****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Recording duration, in milliseconds.
	//
	// example:
	//
	// 56321
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Recording file name.
	//
	// example:
	//
	// job-25920271311543****-798f1e90-1f82-42da-914c-46580c8f4c85-1656926518491.mkv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// OSS download URL for the recording file. Note the time-to-live (TTL) of the download URL. The download URL is valid for 1 day.
	//
	// example:
	//
	// https://ccc-v2-shanghai.oss-cn-shanghai.aliyuncs.com/ccc-test/job-25920271311543****-798f1e90-1f82-42da-914c-46580c8f4c85-1656926518491.mkv?Expires=1657014031&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// List of call hold time segments.
	HoldTimeSegments []*ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments `json:"HoldTimeSegments,omitempty" xml:"HoldTimeSegments,omitempty" type:"Repeated"`
	// RAM account ID for the agent.
	//
	// example:
	//
	// 22807673106369****
	RamId *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// Skill group ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Recording start time, in UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1656926518491
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMultiChannelRecordingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMultiChannelRecordingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetAgentChannelId() *string {
	return s.AgentChannelId
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetDuration() *string {
	return s.Duration
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetHoldTimeSegments() []*ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments {
	return s.HoldTimeSegments
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetRamId() *string {
	return s.RamId
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListMultiChannelRecordingsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetAgentChannelId(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.AgentChannelId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetAgentId(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetAgentName(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetContactId(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetDuration(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetFileName(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetFileUrl(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetHoldTimeSegments(v []*ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) *ListMultiChannelRecordingsResponseBodyData {
	s.HoldTimeSegments = v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetRamId(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetSkillGroupId(v string) *ListMultiChannelRecordingsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) SetStartTime(v int64) *ListMultiChannelRecordingsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyData) Validate() error {
	if s.HoldTimeSegments != nil {
		for _, item := range s.HoldTimeSegments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments struct {
	// The end time of the call hold, in milliseconds.
	//
	// example:
	//
	// 1687860143925
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Call hold start time, in milliseconds.
	//
	// example:
	//
	// 1673255098049
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) String() string {
	return dara.Prettify(s)
}

func (s ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) GoString() string {
	return s.String()
}

func (s *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) SetEndTime(v int64) *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments {
	s.EndTime = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) SetStartTime(v int64) *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments {
	s.StartTime = &v
	return s
}

func (s *ListMultiChannelRecordingsResponseBodyDataHoldTimeSegments) Validate() error {
	return dara.Validate(s)
}
