// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniRealtimeConversationEUResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OmniRealtimeConversationEUResponseBody
	GetCode() *string
	SetData(v *OmniRealtimeConversationEUResponseBodyData) *OmniRealtimeConversationEUResponseBody
	GetData() *OmniRealtimeConversationEUResponseBodyData
	SetHttpStatusCode(v int32) *OmniRealtimeConversationEUResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OmniRealtimeConversationEUResponseBody
	GetMessage() *string
	SetRequestId(v string) *OmniRealtimeConversationEUResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OmniRealtimeConversationEUResponseBody
	GetSuccess() *bool
}

type OmniRealtimeConversationEUResponseBody struct {
	// example:
	//
	// success
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *OmniRealtimeConversationEUResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 89946BAA-E4E1-5114-8A5E-A542FEDC9B16
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OmniRealtimeConversationEUResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OmniRealtimeConversationEUResponseBody) GoString() string {
	return s.String()
}

func (s *OmniRealtimeConversationEUResponseBody) GetCode() *string {
	return s.Code
}

func (s *OmniRealtimeConversationEUResponseBody) GetData() *OmniRealtimeConversationEUResponseBodyData {
	return s.Data
}

func (s *OmniRealtimeConversationEUResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OmniRealtimeConversationEUResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OmniRealtimeConversationEUResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OmniRealtimeConversationEUResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OmniRealtimeConversationEUResponseBody) SetCode(v string) *OmniRealtimeConversationEUResponseBody {
	s.Code = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) SetData(v *OmniRealtimeConversationEUResponseBodyData) *OmniRealtimeConversationEUResponseBody {
	s.Data = v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) SetHttpStatusCode(v int32) *OmniRealtimeConversationEUResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) SetMessage(v string) *OmniRealtimeConversationEUResponseBody {
	s.Message = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) SetRequestId(v string) *OmniRealtimeConversationEUResponseBody {
	s.RequestId = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) SetSuccess(v bool) *OmniRealtimeConversationEUResponseBody {
	s.Success = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OmniRealtimeConversationEUResponseBodyData struct {
	AudioResult *string `json:"audioResult,omitempty" xml:"audioResult,omitempty"`
}

func (s OmniRealtimeConversationEUResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OmniRealtimeConversationEUResponseBodyData) GoString() string {
	return s.String()
}

func (s *OmniRealtimeConversationEUResponseBodyData) GetAudioResult() *string {
	return s.AudioResult
}

func (s *OmniRealtimeConversationEUResponseBodyData) SetAudioResult(v string) *OmniRealtimeConversationEUResponseBodyData {
	s.AudioResult = &v
	return s
}

func (s *OmniRealtimeConversationEUResponseBodyData) Validate() error {
	return dara.Validate(s)
}
