// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoiceCloneResponseBody
	GetCode() *string
	SetData(v *ListVoiceCloneResponseBodyData) *ListVoiceCloneResponseBody
	GetData() *ListVoiceCloneResponseBodyData
	SetHttpStatusCode(v int32) *ListVoiceCloneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoiceCloneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVoiceCloneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVoiceCloneResponseBody
	GetSuccess() *bool
}

type ListVoiceCloneResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListVoiceCloneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A573837-3AD3-54CF-930A-07A3287042C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVoiceCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceCloneResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceCloneResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoiceCloneResponseBody) GetData() *ListVoiceCloneResponseBodyData {
	return s.Data
}

func (s *ListVoiceCloneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoiceCloneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoiceCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceCloneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVoiceCloneResponseBody) SetCode(v string) *ListVoiceCloneResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceCloneResponseBody) SetData(v *ListVoiceCloneResponseBodyData) *ListVoiceCloneResponseBody {
	s.Data = v
	return s
}

func (s *ListVoiceCloneResponseBody) SetHttpStatusCode(v int32) *ListVoiceCloneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoiceCloneResponseBody) SetMessage(v string) *ListVoiceCloneResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceCloneResponseBody) SetRequestId(v string) *ListVoiceCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceCloneResponseBody) SetSuccess(v bool) *ListVoiceCloneResponseBody {
	s.Success = &v
	return s
}

func (s *ListVoiceCloneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceCloneResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of cloned voices that meet the filter conditions.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of cloned voices.
	VoiceClones []*ListVoiceCloneResponseBodyDataVoiceClones `json:"VoiceClones,omitempty" xml:"VoiceClones,omitempty" type:"Repeated"`
}

func (s ListVoiceCloneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceCloneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoiceCloneResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceCloneResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceCloneResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoiceCloneResponseBodyData) GetVoiceClones() []*ListVoiceCloneResponseBodyDataVoiceClones {
	return s.VoiceClones
}

func (s *ListVoiceCloneResponseBodyData) SetPageNumber(v int32) *ListVoiceCloneResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceCloneResponseBodyData) SetPageSize(v int32) *ListVoiceCloneResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoiceCloneResponseBodyData) SetTotalCount(v int32) *ListVoiceCloneResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoiceCloneResponseBodyData) SetVoiceClones(v []*ListVoiceCloneResponseBodyDataVoiceClones) *ListVoiceCloneResponseBodyData {
	s.VoiceClones = v
	return s
}

func (s *ListVoiceCloneResponseBodyData) Validate() error {
	if s.VoiceClones != nil {
		for _, item := range s.VoiceClones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVoiceCloneResponseBodyDataVoiceClones struct {
	// The Alibaba Cloud account ID of the creator of the voice.
	//
	// example:
	//
	// 1353945678901234
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The type of the text-to-speech (TTS) engine used by the voice.
	//
	// example:
	//
	// bailian
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1791891200000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1791891200000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The identifier of the model used for cloning.
	//
	// example:
	//
	// cosyvoice-v3.5-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The voice status. Valid values: Enabled (available) and Disabled (unavailable).
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// cosyvoice-v2-067dc4b2e9d34dbeb5de20b5b093d1e3
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	//
	// example:
	//
	// Cloned Voice 1
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s ListVoiceCloneResponseBodyDataVoiceClones) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceCloneResponseBodyDataVoiceClones) GoString() string {
	return s.String()
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetCreator() *string {
	return s.Creator
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetEngine() *string {
	return s.Engine
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetModel() *string {
	return s.Model
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetStatus() *string {
	return s.Status
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetVoiceId() *string {
	return s.VoiceId
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) GetVoiceName() *string {
	return s.VoiceName
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetCreator(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.Creator = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetEngine(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.Engine = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetGmtCreate(v int64) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.GmtCreate = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetGmtModified(v int64) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.GmtModified = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetInstanceId(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.InstanceId = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetModel(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.Model = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetStatus(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.Status = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetVoiceId(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.VoiceId = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) SetVoiceName(v string) *ListVoiceCloneResponseBodyDataVoiceClones {
	s.VoiceName = &v
	return s
}

func (s *ListVoiceCloneResponseBodyDataVoiceClones) Validate() error {
	return dara.Validate(s)
}
