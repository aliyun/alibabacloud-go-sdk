// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramTypeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramTypeAndCounts(v []*ListProgramTypeCountResponseBodyProgramTypeAndCounts) *ListProgramTypeCountResponseBody
	GetProgramTypeAndCounts() []*ListProgramTypeCountResponseBodyProgramTypeAndCounts
	SetRequestId(v string) *ListProgramTypeCountResponseBody
	GetRequestId() *string
}

type ListProgramTypeCountResponseBody struct {
	// The list of node types and quantity.
	ProgramTypeAndCounts []*ListProgramTypeCountResponseBodyProgramTypeAndCounts `json:"ProgramTypeAndCounts,omitempty" xml:"ProgramTypeAndCounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProgramTypeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProgramTypeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ListProgramTypeCountResponseBody) GetProgramTypeAndCounts() []*ListProgramTypeCountResponseBodyProgramTypeAndCounts {
	return s.ProgramTypeAndCounts
}

func (s *ListProgramTypeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProgramTypeCountResponseBody) SetProgramTypeAndCounts(v []*ListProgramTypeCountResponseBodyProgramTypeAndCounts) *ListProgramTypeCountResponseBody {
	s.ProgramTypeAndCounts = v
	return s
}

func (s *ListProgramTypeCountResponseBody) SetRequestId(v string) *ListProgramTypeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProgramTypeCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProgramTypeCountResponseBodyProgramTypeAndCounts struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The node type.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
}

func (s ListProgramTypeCountResponseBodyProgramTypeAndCounts) String() string {
	return dara.Prettify(s)
}

func (s ListProgramTypeCountResponseBodyProgramTypeAndCounts) GoString() string {
	return s.String()
}

func (s *ListProgramTypeCountResponseBodyProgramTypeAndCounts) GetCount() *int32 {
	return s.Count
}

func (s *ListProgramTypeCountResponseBodyProgramTypeAndCounts) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListProgramTypeCountResponseBodyProgramTypeAndCounts) SetCount(v int32) *ListProgramTypeCountResponseBodyProgramTypeAndCounts {
	s.Count = &v
	return s
}

func (s *ListProgramTypeCountResponseBodyProgramTypeAndCounts) SetProgramType(v string) *ListProgramTypeCountResponseBodyProgramTypeAndCounts {
	s.ProgramType = &v
	return s
}

func (s *ListProgramTypeCountResponseBodyProgramTypeAndCounts) Validate() error {
	return dara.Validate(s)
}
