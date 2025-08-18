// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRoutineUserInfoResponseBody
	GetRequestId() *string
	SetRoutines(v []*GetRoutineUserInfoResponseBodyRoutines) *GetRoutineUserInfoResponseBody
	GetRoutines() []*GetRoutineUserInfoResponseBodyRoutines
	SetSubdomains(v []*string) *GetRoutineUserInfoResponseBody
	GetSubdomains() []*string
}

type GetRoutineUserInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routines.
	Routines []*GetRoutineUserInfoResponseBodyRoutines `json:"Routines,omitempty" xml:"Routines,omitempty" type:"Repeated"`
	// The subdomains.
	Subdomains []*string `json:"Subdomains,omitempty" xml:"Subdomains,omitempty" type:"Repeated"`
}

func (s GetRoutineUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineUserInfoResponseBody) GetRoutines() []*GetRoutineUserInfoResponseBodyRoutines {
	return s.Routines
}

func (s *GetRoutineUserInfoResponseBody) GetSubdomains() []*string {
	return s.Subdomains
}

func (s *GetRoutineUserInfoResponseBody) SetRequestId(v string) *GetRoutineUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineUserInfoResponseBody) SetRoutines(v []*GetRoutineUserInfoResponseBodyRoutines) *GetRoutineUserInfoResponseBody {
	s.Routines = v
	return s
}

func (s *GetRoutineUserInfoResponseBody) SetSubdomains(v []*string) *GetRoutineUserInfoResponseBody {
	s.Subdomains = v
	return s
}

func (s *GetRoutineUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRoutineUserInfoResponseBodyRoutines struct {
	// The time when the routine was created.
	//
	// example:
	//
	// 2024-03-11T01:23:02.883361712Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The routine description, which is Base64-encoded.
	//
	// example:
	//
	// ZWRpdCByb3V0aW5lIGNvbmZpZyBkZXNjcmlwdGlvbg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The routine name.
	//
	// example:
	//
	// test-routine1
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s GetRoutineUserInfoResponseBodyRoutines) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineUserInfoResponseBodyRoutines) GoString() string {
	return s.String()
}

func (s *GetRoutineUserInfoResponseBodyRoutines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineUserInfoResponseBodyRoutines) GetDescription() *string {
	return s.Description
}

func (s *GetRoutineUserInfoResponseBodyRoutines) GetRoutineName() *string {
	return s.RoutineName
}

func (s *GetRoutineUserInfoResponseBodyRoutines) SetCreateTime(v string) *GetRoutineUserInfoResponseBodyRoutines {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineUserInfoResponseBodyRoutines) SetDescription(v string) *GetRoutineUserInfoResponseBodyRoutines {
	s.Description = &v
	return s
}

func (s *GetRoutineUserInfoResponseBodyRoutines) SetRoutineName(v string) *GetRoutineUserInfoResponseBodyRoutines {
	s.RoutineName = &v
	return s
}

func (s *GetRoutineUserInfoResponseBodyRoutines) Validate() error {
	return dara.Validate(s)
}
