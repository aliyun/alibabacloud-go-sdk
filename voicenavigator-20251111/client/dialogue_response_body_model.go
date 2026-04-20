// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DialogueResponseBody
	GetCode() *string
	SetData(v *DialogueResponseBodyData) *DialogueResponseBody
	GetData() *DialogueResponseBodyData
	SetHttpStatusCode(v int32) *DialogueResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DialogueResponseBody
	GetMessage() *string
	SetParams(v []*string) *DialogueResponseBody
	GetParams() []*string
	SetRequestId(v string) *DialogueResponseBody
	GetRequestId() *string
}

type DialogueResponseBody struct {
	// example:
	//
	// OK
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DialogueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DialogueResponseBody) GetCode() *string {
	return s.Code
}

func (s *DialogueResponseBody) GetData() *DialogueResponseBodyData {
	return s.Data
}

func (s *DialogueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DialogueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DialogueResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DialogueResponseBody) SetCode(v string) *DialogueResponseBody {
	s.Code = &v
	return s
}

func (s *DialogueResponseBody) SetData(v *DialogueResponseBodyData) *DialogueResponseBody {
	s.Data = v
	return s
}

func (s *DialogueResponseBody) SetHttpStatusCode(v int32) *DialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DialogueResponseBody) SetMessage(v string) *DialogueResponseBody {
	s.Message = &v
	return s
}

func (s *DialogueResponseBody) SetParams(v []*string) *DialogueResponseBody {
	s.Params = v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DialogueResponseBodyData struct {
	Answer            *string                                      `json:"Answer,omitempty" xml:"Answer,omitempty"`
	ControlParamsList []*DialogueResponseBodyDataControlParamsList `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty" type:"Repeated"`
	// example:
	//
	// 1752283603978
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 07d72ea0-6e38-48d4-8371-14deaaba996f
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1774858849987
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// true
	StreamEnd *bool `json:"StreamEnd,omitempty" xml:"StreamEnd,omitempty"`
	// example:
	//
	// 1
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s DialogueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *DialogueResponseBodyData) GetControlParamsList() []*DialogueResponseBodyDataControlParamsList {
	return s.ControlParamsList
}

func (s *DialogueResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DialogueResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DialogueResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DialogueResponseBodyData) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *DialogueResponseBodyData) GetStreamId() *string {
	return s.StreamId
}

func (s *DialogueResponseBodyData) SetAnswer(v string) *DialogueResponseBodyData {
	s.Answer = &v
	return s
}

func (s *DialogueResponseBodyData) SetControlParamsList(v []*DialogueResponseBodyDataControlParamsList) *DialogueResponseBodyData {
	s.ControlParamsList = v
	return s
}

func (s *DialogueResponseBodyData) SetEndTime(v int64) *DialogueResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DialogueResponseBodyData) SetSessionId(v string) *DialogueResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DialogueResponseBodyData) SetStartTime(v int64) *DialogueResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DialogueResponseBodyData) SetStreamEnd(v bool) *DialogueResponseBodyData {
	s.StreamEnd = &v
	return s
}

func (s *DialogueResponseBodyData) SetStreamId(v string) *DialogueResponseBodyData {
	s.StreamId = &v
	return s
}

func (s *DialogueResponseBodyData) Validate() error {
	if s.ControlParamsList != nil {
		for _, item := range s.ControlParamsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DialogueResponseBodyDataControlParamsList struct {
	// example:
	//
	// Voice
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DialogueResponseBodyDataControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyDataControlParamsList) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyDataControlParamsList) GetType() *string {
	return s.Type
}

func (s *DialogueResponseBodyDataControlParamsList) SetType(v string) *DialogueResponseBodyDataControlParamsList {
	s.Type = &v
	return s
}

func (s *DialogueResponseBodyDataControlParamsList) Validate() error {
	return dara.Validate(s)
}
