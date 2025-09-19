// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotEventTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotEventTrendResponseBody
	GetCode() *string
	SetCount(v int32) *GetHoneypotEventTrendResponseBody
	GetCount() *int32
	SetData(v []*GetHoneypotEventTrendResponseBodyData) *GetHoneypotEventTrendResponseBody
	GetData() []*GetHoneypotEventTrendResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneypotEventTrendResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotEventTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotEventTrendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotEventTrendResponseBody
	GetSuccess() *bool
}

type GetHoneypotEventTrendResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of data returned. ￼
	Data []*GetHoneypotEventTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 21637690-9B21-5EEC-94DB-2A732480****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneypotEventTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotEventTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotEventTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotEventTrendResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetHoneypotEventTrendResponseBody) GetData() []*GetHoneypotEventTrendResponseBodyData {
	return s.Data
}

func (s *GetHoneypotEventTrendResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotEventTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotEventTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotEventTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotEventTrendResponseBody) SetCode(v string) *GetHoneypotEventTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetCount(v int32) *GetHoneypotEventTrendResponseBody {
	s.Count = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetData(v []*GetHoneypotEventTrendResponseBodyData) *GetHoneypotEventTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetHttpStatusCode(v int32) *GetHoneypotEventTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetMessage(v string) *GetHoneypotEventTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetRequestId(v string) *GetHoneypotEventTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) SetSuccess(v bool) *GetHoneypotEventTrendResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotEventTrendResponseBodyData struct {
	// The number of attacks that occurred.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the attack.
	//
	// example:
	//
	// ATTACK_EVENT
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The timestamp when the attack event occurred.
	//
	// example:
	//
	// 1686968163644
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s GetHoneypotEventTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotEventTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneypotEventTrendResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetHoneypotEventTrendResponseBodyData) GetKeyName() *string {
	return s.KeyName
}

func (s *GetHoneypotEventTrendResponseBodyData) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *GetHoneypotEventTrendResponseBodyData) SetCount(v int32) *GetHoneypotEventTrendResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBodyData) SetKeyName(v string) *GetHoneypotEventTrendResponseBodyData {
	s.KeyName = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBodyData) SetTimeStamp(v int64) *GetHoneypotEventTrendResponseBodyData {
	s.TimeStamp = &v
	return s
}

func (s *GetHoneypotEventTrendResponseBodyData) Validate() error {
	return dara.Validate(s)
}
