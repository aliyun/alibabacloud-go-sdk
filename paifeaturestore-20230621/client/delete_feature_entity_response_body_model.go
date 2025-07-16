// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFeatureEntityResponseBody
	GetRequestId() *string
}

type DeleteFeatureEntityResponseBody struct {
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFeatureEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFeatureEntityResponseBody) SetRequestId(v string) *DeleteFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFeatureEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
