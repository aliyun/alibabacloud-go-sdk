// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeFreeUserEventTypesResponseBodyData) *DescribeFreeUserEventTypesResponseBody
	GetData() []*DescribeFreeUserEventTypesResponseBodyData
	SetRequestId(v string) *DescribeFreeUserEventTypesResponseBody
	GetRequestId() *string
}

type DescribeFreeUserEventTypesResponseBody struct {
	// The types of security events on which basic detection is performed.
	Data []*DescribeFreeUserEventTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B9D6AD11-DD3D-5A27-B1D9-8A37F7777196
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFreeUserEventTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventTypesResponseBody) GetData() []*DescribeFreeUserEventTypesResponseBodyData {
	return s.Data
}

func (s *DescribeFreeUserEventTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFreeUserEventTypesResponseBody) SetData(v []*DescribeFreeUserEventTypesResponseBodyData) *DescribeFreeUserEventTypesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFreeUserEventTypesResponseBody) SetRequestId(v string) *DescribeFreeUserEventTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFreeUserEventTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFreeUserEventTypesResponseBodyData struct {
	// The number of security events.
	//
	// example:
	//
	// 4
	EventNum *string `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// The type of the security event.
	//
	// example:
	//
	// SMSInterfaceAbuse
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s DescribeFreeUserEventTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventTypesResponseBodyData) GetEventNum() *string {
	return s.EventNum
}

func (s *DescribeFreeUserEventTypesResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *DescribeFreeUserEventTypesResponseBodyData) SetEventNum(v string) *DescribeFreeUserEventTypesResponseBodyData {
	s.EventNum = &v
	return s
}

func (s *DescribeFreeUserEventTypesResponseBodyData) SetEventType(v string) *DescribeFreeUserEventTypesResponseBodyData {
	s.EventType = &v
	return s
}

func (s *DescribeFreeUserEventTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
