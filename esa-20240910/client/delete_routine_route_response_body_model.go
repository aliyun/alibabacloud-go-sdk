// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoutineRouteResponseBody
	GetRequestId() *string
}

type DeleteRoutineRouteResponseBody struct {
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineRouteResponseBody) SetRequestId(v string) *DeleteRoutineRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
