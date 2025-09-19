// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNsasSuspEventTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypes(v []*DescribeNsasSuspEventTypeResponseBodyEventTypes) *DescribeNsasSuspEventTypeResponseBody
	GetEventTypes() []*DescribeNsasSuspEventTypeResponseBodyEventTypes
	SetRequestId(v string) *DescribeNsasSuspEventTypeResponseBody
	GetRequestId() *string
}

type DescribeNsasSuspEventTypeResponseBody struct {
	// An array that consists of the information about the alert type.
	EventTypes []*DescribeNsasSuspEventTypeResponseBodyEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9FBC6E47-7508-58C9-9E76-528E118CB1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNsasSuspEventTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNsasSuspEventTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNsasSuspEventTypeResponseBody) GetEventTypes() []*DescribeNsasSuspEventTypeResponseBodyEventTypes {
	return s.EventTypes
}

func (s *DescribeNsasSuspEventTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNsasSuspEventTypeResponseBody) SetEventTypes(v []*DescribeNsasSuspEventTypeResponseBodyEventTypes) *DescribeNsasSuspEventTypeResponseBody {
	s.EventTypes = v
	return s
}

func (s *DescribeNsasSuspEventTypeResponseBody) SetRequestId(v string) *DescribeNsasSuspEventTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNsasSuspEventTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNsasSuspEventTypeResponseBodyEventTypes struct {
	// The name of the alert type.
	//
	// example:
	//
	// Unusual Logon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of assets for which an alert of the type is generated.
	//
	// example:
	//
	// 22
	SuspEventCount *int32 `json:"SuspEventCount,omitempty" xml:"SuspEventCount,omitempty"`
	// The alert type.
	//
	// example:
	//
	// Unusual Logon
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNsasSuspEventTypeResponseBodyEventTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeNsasSuspEventTypeResponseBodyEventTypes) GoString() string {
	return s.String()
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) GetName() *string {
	return s.Name
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) GetSuspEventCount() *int32 {
	return s.SuspEventCount
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) GetType() *string {
	return s.Type
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) SetName(v string) *DescribeNsasSuspEventTypeResponseBodyEventTypes {
	s.Name = &v
	return s
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) SetSuspEventCount(v int32) *DescribeNsasSuspEventTypeResponseBodyEventTypes {
	s.SuspEventCount = &v
	return s
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) SetType(v string) *DescribeNsasSuspEventTypeResponseBodyEventTypes {
	s.Type = &v
	return s
}

func (s *DescribeNsasSuspEventTypeResponseBodyEventTypes) Validate() error {
	return dara.Validate(s)
}
