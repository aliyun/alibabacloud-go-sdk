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
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Routines   []*GetRoutineUserInfoResponseBodyRoutines `json:"Routines,omitempty" xml:"Routines,omitempty" type:"Repeated"`
	Subdomains []*string                                 `json:"Subdomains,omitempty" xml:"Subdomains,omitempty" type:"Repeated"`
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
	if s.Routines != nil {
		for _, item := range s.Routines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRoutineUserInfoResponseBodyRoutines struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
