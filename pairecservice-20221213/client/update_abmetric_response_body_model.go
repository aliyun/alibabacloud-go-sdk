// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABMetricResponseBody
	GetRequestId() *string
}

type UpdateABMetricResponseBody struct {
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateABMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABMetricResponseBody) SetRequestId(v string) *UpdateABMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
