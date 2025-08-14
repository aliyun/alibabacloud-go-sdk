// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenInterRegionTrafficQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenInterRegionTrafficQosPolicyResponseBody
	GetRequestId() *string
}

type DeleteCenInterRegionTrafficQosPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6DF9A765-BCD2-5C7E-8C32-C35C8A361A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponseBody) SetRequestId(v string) *DeleteCenInterRegionTrafficQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
