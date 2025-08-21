// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActionPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteActionPlanResponseBody
	GetRequestId() *string
}

type DeleteActionPlanResponseBody struct {
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteActionPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteActionPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteActionPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteActionPlanResponseBody) SetRequestId(v string) *DeleteActionPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteActionPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
