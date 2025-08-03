// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbnormalEventsCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAbnormalEventsCountResponseBody
	GetCode() *string
	SetData(v []*GetAbnormalEventsCountResponseBodyData) *GetAbnormalEventsCountResponseBody
	GetData() []*GetAbnormalEventsCountResponseBodyData
	SetMessage(v string) *GetAbnormalEventsCountResponseBody
	GetMessage() *string
}

type GetAbnormalEventsCountResponseBody struct {
	// example:
	//
	// Success
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetAbnormalEventsCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAbnormalEventsCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAbnormalEventsCountResponseBody) GetData() []*GetAbnormalEventsCountResponseBodyData {
	return s.Data
}

func (s *GetAbnormalEventsCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAbnormalEventsCountResponseBody) SetCode(v string) *GetAbnormalEventsCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetData(v []*GetAbnormalEventsCountResponseBodyData) *GetAbnormalEventsCountResponseBody {
	s.Data = v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetMessage(v string) *GetAbnormalEventsCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAbnormalEventsCountResponseBodyData struct {
	EventList []*string `json:"eventList,omitempty" xml:"eventList,omitempty" type:"Repeated"`
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAbnormalEventsCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBodyData) GetEventList() []*string {
	return s.EventList
}

func (s *GetAbnormalEventsCountResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetAbnormalEventsCountResponseBodyData) GetValue() *int64 {
	return s.Value
}

func (s *GetAbnormalEventsCountResponseBodyData) SetEventList(v []*string) *GetAbnormalEventsCountResponseBodyData {
	s.EventList = v
	return s
}

func (s *GetAbnormalEventsCountResponseBodyData) SetType(v string) *GetAbnormalEventsCountResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBodyData) SetValue(v int64) *GetAbnormalEventsCountResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
