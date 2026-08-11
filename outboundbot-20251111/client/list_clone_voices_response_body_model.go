// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCloneVoicesResponseBody
	GetCode() *string
	SetData(v *ListCloneVoicesResponseBodyData) *ListCloneVoicesResponseBody
	GetData() *ListCloneVoicesResponseBodyData
	SetHttpStatusCode(v int32) *ListCloneVoicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCloneVoicesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCloneVoicesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCloneVoicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloneVoicesResponseBody
	GetSuccess() *bool
}

type ListCloneVoicesResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListCloneVoicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=anchashi.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloneVoicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloneVoicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCloneVoicesResponseBody) GetData() *ListCloneVoicesResponseBodyData {
	return s.Data
}

func (s *ListCloneVoicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCloneVoicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloneVoicesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCloneVoicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloneVoicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloneVoicesResponseBody) SetCode(v string) *ListCloneVoicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloneVoicesResponseBody) SetData(v *ListCloneVoicesResponseBodyData) *ListCloneVoicesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloneVoicesResponseBody) SetHttpStatusCode(v int32) *ListCloneVoicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCloneVoicesResponseBody) SetMessage(v string) *ListCloneVoicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloneVoicesResponseBody) SetParams(v []*string) *ListCloneVoicesResponseBody {
	s.Params = v
	return s
}

func (s *ListCloneVoicesResponseBody) SetRequestId(v string) *ListCloneVoicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloneVoicesResponseBody) SetSuccess(v bool) *ListCloneVoicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloneVoicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloneVoicesResponseBodyData struct {
	// The list of cloned voices.
	CloneVoices []*ListCloneVoicesResponseBodyDataCloneVoices `json:"CloneVoices,omitempty" xml:"CloneVoices,omitempty" type:"Repeated"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloneVoicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloneVoicesResponseBodyData) GetCloneVoices() []*ListCloneVoicesResponseBodyDataCloneVoices {
	return s.CloneVoices
}

func (s *ListCloneVoicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloneVoicesResponseBodyData) SetCloneVoices(v []*ListCloneVoicesResponseBodyDataCloneVoices) *ListCloneVoicesResponseBodyData {
	s.CloneVoices = v
	return s
}

func (s *ListCloneVoicesResponseBodyData) SetPageNumber(v int32) *ListCloneVoicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoicesResponseBodyData) SetPageSize(v int32) *ListCloneVoicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoicesResponseBodyData) SetTotalCount(v int32) *ListCloneVoicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCloneVoicesResponseBodyData) Validate() error {
	if s.CloneVoices != nil {
		for _, item := range s.CloneVoices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloneVoicesResponseBodyDataCloneVoices struct {
	// The UUID of the cloned voice.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model name.
	//
	// example:
	//
	// CosyVoice
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name.
	//
	// example:
	//
	// TestClonedVoice
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The speech vendor.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The status.
	//
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1308144684576765
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// cosyvoice-v3-flash-voicebot2-8aa485413eba42089c873eec1f901d64
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s ListCloneVoicesResponseBodyDataCloneVoices) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoicesResponseBodyDataCloneVoices) GoString() string {
	return s.String()
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetModel() *string {
	return s.Model
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetName() *string {
	return s.Name
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetStatus() *string {
	return s.Status
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) GetVoice() *string {
	return s.Voice
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetCloneVoiceId(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.CloneVoiceId = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetCreatedTime(v int64) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.CreatedTime = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetInstanceId(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.InstanceId = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetModel(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.Model = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetName(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.Name = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetNlsEngine(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.NlsEngine = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetStatus(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.Status = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetTenantId(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.TenantId = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetUpdatedTime(v int64) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.UpdatedTime = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) SetVoice(v string) *ListCloneVoicesResponseBodyDataCloneVoices {
	s.Voice = &v
	return s
}

func (s *ListCloneVoicesResponseBodyDataCloneVoices) Validate() error {
	return dara.Validate(s)
}
