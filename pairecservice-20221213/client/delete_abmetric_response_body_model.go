// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteABMetricResponseBody
	GetRequestId() *string
}

type DeleteABMetricResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteABMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteABMetricResponseBody) SetRequestId(v string) *DeleteABMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
