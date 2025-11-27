// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateActionPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateActionPlanResponseBody
	GetRequestId() *string
}

type UpdateActionPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateActionPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateActionPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateActionPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateActionPlanResponseBody) SetRequestId(v string) *UpdateActionPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateActionPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
