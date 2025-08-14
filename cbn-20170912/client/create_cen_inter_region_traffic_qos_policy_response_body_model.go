// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody
	GetRequestId() *string
	SetTrafficQosPolicyId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody
	GetTrafficQosPolicyId() *string
}

type CreateCenInterRegionTrafficQosPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6DF9A765-BCD2-5C7E-8C32-C35C8A361A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-eczzew0v1kzrb5****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) SetRequestId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) SetTrafficQosPolicyId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
