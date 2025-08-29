// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteABMetricGroupResponseBody
	GetRequestId() *string
}

type DeleteABMetricGroupResponseBody struct {
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteABMetricGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteABMetricGroupResponseBody) SetRequestId(v string) *DeleteABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABMetricGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
