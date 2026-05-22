// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCanaryAreasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanaryAreas(v []*string) *ListRoutineCanaryAreasResponseBody
	GetCanaryAreas() []*string
	SetRequestId(v string) *ListRoutineCanaryAreasResponseBody
	GetRequestId() *string
}

type ListRoutineCanaryAreasResponseBody struct {
	// The regions for canary release.
	CanaryAreas []*string `json:"CanaryAreas,omitempty" xml:"CanaryAreas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRoutineCanaryAreasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCanaryAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineCanaryAreasResponseBody) GetCanaryAreas() []*string {
	return s.CanaryAreas
}

func (s *ListRoutineCanaryAreasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineCanaryAreasResponseBody) SetCanaryAreas(v []*string) *ListRoutineCanaryAreasResponseBody {
	s.CanaryAreas = v
	return s
}

func (s *ListRoutineCanaryAreasResponseBody) SetRequestId(v string) *ListRoutineCanaryAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineCanaryAreasResponseBody) Validate() error {
	return dara.Validate(s)
}
