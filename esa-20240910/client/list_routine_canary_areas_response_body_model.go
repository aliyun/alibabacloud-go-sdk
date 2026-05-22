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
	CanaryAreas []*string `json:"CanaryAreas,omitempty" xml:"CanaryAreas,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
