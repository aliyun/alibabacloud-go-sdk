// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRoutineRouteResponseBody
	GetRequestId() *string
}

type UpdateRoutineRouteResponseBody struct {
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRoutineRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoutineRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRoutineRouteResponseBody) SetRequestId(v string) *UpdateRoutineRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoutineRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
