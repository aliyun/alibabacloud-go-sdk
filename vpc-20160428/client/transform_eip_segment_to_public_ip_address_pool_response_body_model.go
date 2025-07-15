// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformEipSegmentToPublicIpAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicIpAddressPoolId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody
	GetPublicIpAddressPoolId() *string
	SetRequestId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody
	GetResourceGroupId() *string
}

type TransformEipSegmentToPublicIpAddressPoolResponseBody struct {
	// The ID of the IP address pool.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the IP address pool belongs.
	//
	// example:
	//
	// rg-acfmxazb4pcdvf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s TransformEipSegmentToPublicIpAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformEipSegmentToPublicIpAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) SetPublicIpAddressPoolId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) SetRequestId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) SetResourceGroupId(v string) *TransformEipSegmentToPublicIpAddressPoolResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
