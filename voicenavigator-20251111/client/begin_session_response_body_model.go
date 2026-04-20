// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BeginSessionResponseBody
	GetCode() *string
	SetData(v *BeginSessionResponseBodyData) *BeginSessionResponseBody
	GetData() *BeginSessionResponseBodyData
	SetHttpStatusCode(v int32) *BeginSessionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BeginSessionResponseBody
	GetMessage() *string
	SetParams(v []*string) *BeginSessionResponseBody
	GetParams() []*string
	SetRequestId(v string) *BeginSessionResponseBody
	GetRequestId() *string
}

type BeginSessionResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BeginSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeginSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *BeginSessionResponseBody) GetData() *BeginSessionResponseBodyData {
	return s.Data
}

func (s *BeginSessionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BeginSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BeginSessionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BeginSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BeginSessionResponseBody) SetCode(v string) *BeginSessionResponseBody {
	s.Code = &v
	return s
}

func (s *BeginSessionResponseBody) SetData(v *BeginSessionResponseBodyData) *BeginSessionResponseBody {
	s.Data = v
	return s
}

func (s *BeginSessionResponseBody) SetHttpStatusCode(v int32) *BeginSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BeginSessionResponseBody) SetMessage(v string) *BeginSessionResponseBody {
	s.Message = &v
	return s
}

func (s *BeginSessionResponseBody) SetParams(v []*string) *BeginSessionResponseBody {
	s.Params = v
	return s
}

func (s *BeginSessionResponseBody) SetRequestId(v string) *BeginSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BeginSessionResponseBodyData struct {
	Answer            *string                                          `json:"Answer,omitempty" xml:"Answer,omitempty"`
	ControlParamsList []*BeginSessionResponseBodyDataControlParamsList `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty" type:"Repeated"`
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

func (s BeginSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeginSessionResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *BeginSessionResponseBodyData) GetControlParamsList() []*BeginSessionResponseBodyDataControlParamsList {
	return s.ControlParamsList
}

func (s *BeginSessionResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *BeginSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *BeginSessionResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *BeginSessionResponseBodyData) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *BeginSessionResponseBodyData) GetStreamId() *string {
	return s.StreamId
}

func (s *BeginSessionResponseBodyData) SetAnswer(v string) *BeginSessionResponseBodyData {
	s.Answer = &v
	return s
}

func (s *BeginSessionResponseBodyData) SetControlParamsList(v []*BeginSessionResponseBodyDataControlParamsList) *BeginSessionResponseBodyData {
	s.ControlParamsList = v
	return s
}

func (s *BeginSessionResponseBodyData) SetEndTime(v int64) *BeginSessionResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *BeginSessionResponseBodyData) SetSessionId(v string) *BeginSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *BeginSessionResponseBodyData) SetStartTime(v int64) *BeginSessionResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *BeginSessionResponseBodyData) SetStreamEnd(v bool) *BeginSessionResponseBodyData {
	s.StreamEnd = &v
	return s
}

func (s *BeginSessionResponseBodyData) SetStreamId(v string) *BeginSessionResponseBodyData {
	s.StreamId = &v
	return s
}

func (s *BeginSessionResponseBodyData) Validate() error {
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

type BeginSessionResponseBodyDataControlParamsList struct {
	// example:
	//
	// Voice
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BeginSessionResponseBodyDataControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionResponseBodyDataControlParamsList) GoString() string {
	return s.String()
}

func (s *BeginSessionResponseBodyDataControlParamsList) GetType() *string {
	return s.Type
}

func (s *BeginSessionResponseBodyDataControlParamsList) SetType(v string) *BeginSessionResponseBodyDataControlParamsList {
	s.Type = &v
	return s
}

func (s *BeginSessionResponseBodyDataControlParamsList) Validate() error {
	return dara.Validate(s)
}
