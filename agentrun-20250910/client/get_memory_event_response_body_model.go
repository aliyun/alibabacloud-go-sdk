// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryEventResponseBody
	GetCode() *string
	SetData(v *GetMemoryEventResponseBodyData) *GetMemoryEventResponseBody
	GetData() *GetMemoryEventResponseBodyData
	SetRequestId(v string) *GetMemoryEventResponseBody
	GetRequestId() *string
}

type GetMemoryEventResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetMemoryEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// BF2A63E4-FCE9-5A65-A60E-4086C8EDBC06
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryEventResponseBody) GetData() *GetMemoryEventResponseBodyData {
	return s.Data
}

func (s *GetMemoryEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryEventResponseBody) SetCode(v string) *GetMemoryEventResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryEventResponseBody) SetData(v *GetMemoryEventResponseBodyData) *GetMemoryEventResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryEventResponseBody) SetRequestId(v string) *GetMemoryEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryEventResponseBodyData struct {
	Event map[string]interface{} `json:"event,omitempty" xml:"event,omitempty"`
}

func (s GetMemoryEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryEventResponseBodyData) GetEvent() map[string]interface{} {
	return s.Event
}

func (s *GetMemoryEventResponseBodyData) SetEvent(v map[string]interface{}) *GetMemoryEventResponseBodyData {
	s.Event = v
	return s
}

func (s *GetMemoryEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
