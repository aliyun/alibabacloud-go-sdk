// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABMetricGroupResponseBody
	GetRequestId() *string
}

type UpdateABMetricGroupResponseBody struct {
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateABMetricGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABMetricGroupResponseBody) SetRequestId(v string) *UpdateABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABMetricGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
