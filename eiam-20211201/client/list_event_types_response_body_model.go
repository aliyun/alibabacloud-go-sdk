// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypes(v []*ListEventTypesResponseBodyEventTypes) *ListEventTypesResponseBody
	GetEventTypes() []*ListEventTypesResponseBodyEventTypes
	SetRequestId(v string) *ListEventTypesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListEventTypesResponseBody
	GetTotalCount() *int64
}

type ListEventTypesResponseBody struct {
	EventTypes []*ListEventTypesResponseBodyEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEventTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventTypesResponseBody) GetEventTypes() []*ListEventTypesResponseBodyEventTypes {
	return s.EventTypes
}

func (s *ListEventTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventTypesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEventTypesResponseBody) SetEventTypes(v []*ListEventTypesResponseBodyEventTypes) *ListEventTypesResponseBody {
	s.EventTypes = v
	return s
}

func (s *ListEventTypesResponseBody) SetRequestId(v string) *ListEventTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventTypesResponseBody) SetTotalCount(v int64) *ListEventTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEventTypesResponseBody) Validate() error {
	if s.EventTypes != nil {
		for _, item := range s.EventTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventTypesResponseBodyEventTypes struct {
	// example:
	//
	// urn:alibaba:idaas:event:user:create
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s ListEventTypesResponseBodyEventTypes) String() string {
	return dara.Prettify(s)
}

func (s ListEventTypesResponseBodyEventTypes) GoString() string {
	return s.String()
}

func (s *ListEventTypesResponseBodyEventTypes) GetEventType() *string {
	return s.EventType
}

func (s *ListEventTypesResponseBodyEventTypes) SetEventType(v string) *ListEventTypesResponseBodyEventTypes {
	s.EventType = &v
	return s
}

func (s *ListEventTypesResponseBodyEventTypes) Validate() error {
	return dara.Validate(s)
}
