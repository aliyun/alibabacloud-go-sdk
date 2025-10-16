// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemorySessionResponseBody
	GetCode() *string
	SetData(v *GetMemorySessionResponseBodyData) *GetMemorySessionResponseBody
	GetData() *GetMemorySessionResponseBodyData
	SetRequestId(v string) *GetMemorySessionResponseBody
	GetRequestId() *string
}

type GetMemorySessionResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetMemorySessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 55D4BE40-2811-5CFB-8482-E0E98D575B1E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemorySessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemorySessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemorySessionResponseBody) GetData() *GetMemorySessionResponseBodyData {
	return s.Data
}

func (s *GetMemorySessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemorySessionResponseBody) SetCode(v string) *GetMemorySessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemorySessionResponseBody) SetData(v *GetMemorySessionResponseBodyData) *GetMemorySessionResponseBody {
	s.Data = v
	return s
}

func (s *GetMemorySessionResponseBody) SetRequestId(v string) *GetMemorySessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemorySessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemorySessionResponseBodyData struct {
	Events []map[string]interface{} `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
}

func (s GetMemorySessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemorySessionResponseBodyData) GetEvents() []map[string]interface{} {
	return s.Events
}

func (s *GetMemorySessionResponseBodyData) SetEvents(v []map[string]interface{}) *GetMemorySessionResponseBodyData {
	s.Events = v
	return s
}

func (s *GetMemorySessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
