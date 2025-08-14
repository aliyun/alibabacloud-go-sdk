// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosPolicyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody
	GetRequestId() *string
}

type UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B8465FF-E697-5D3D-AAD5-0B4EEADFDB27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) SetRequestId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
