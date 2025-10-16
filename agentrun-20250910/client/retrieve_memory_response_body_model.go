// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetrieveMemoryResponseBody
	GetCode() *string
	SetData(v *RetrieveMemoryResponseBodyData) *RetrieveMemoryResponseBody
	GetData() *RetrieveMemoryResponseBodyData
	SetRequestId(v string) *RetrieveMemoryResponseBody
	GetRequestId() *string
}

type RetrieveMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *RetrieveMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// BC65E49E-1F6A-55E0-8A0E-7255DBFAA8F9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RetrieveMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetrieveMemoryResponseBody) GetData() *RetrieveMemoryResponseBodyData {
	return s.Data
}

func (s *RetrieveMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveMemoryResponseBody) SetCode(v string) *RetrieveMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveMemoryResponseBody) SetData(v *RetrieveMemoryResponseBodyData) *RetrieveMemoryResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveMemoryResponseBody) SetRequestId(v string) *RetrieveMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveMemoryResponseBodyData struct {
	Events   []map[string]*string `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	Memories []map[string]*string `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
}

func (s RetrieveMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveMemoryResponseBodyData) GetEvents() []map[string]*string {
	return s.Events
}

func (s *RetrieveMemoryResponseBodyData) GetMemories() []map[string]*string {
	return s.Memories
}

func (s *RetrieveMemoryResponseBodyData) SetEvents(v []map[string]*string) *RetrieveMemoryResponseBodyData {
	s.Events = v
	return s
}

func (s *RetrieveMemoryResponseBodyData) SetMemories(v []map[string]*string) *RetrieveMemoryResponseBodyData {
	s.Memories = v
	return s
}

func (s *RetrieveMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
