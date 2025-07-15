// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublicIpAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicIpAddressPoolId(v string) *CreatePublicIpAddressPoolResponseBody
	GetPublicIpAddressPoolId() *string
	SetPulbicIpAddressPoolId(v string) *CreatePublicIpAddressPoolResponseBody
	GetPulbicIpAddressPoolId() *string
	SetRequestId(v string) *CreatePublicIpAddressPoolResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreatePublicIpAddressPoolResponseBody
	GetResourceGroupId() *string
}

type CreatePublicIpAddressPoolResponseBody struct {
	// The ID of the IP address pool.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// Deprecated
	//
	// The ID of the IP address pool.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PulbicIpAddressPoolId *string `json:"PulbicIpAddressPoolId,omitempty" xml:"PulbicIpAddressPoolId,omitempty"`
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

func (s CreatePublicIpAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicIpAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicIpAddressPoolResponseBody) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *CreatePublicIpAddressPoolResponseBody) GetPulbicIpAddressPoolId() *string {
	return s.PulbicIpAddressPoolId
}

func (s *CreatePublicIpAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePublicIpAddressPoolResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePublicIpAddressPoolResponseBody) SetPublicIpAddressPoolId(v string) *CreatePublicIpAddressPoolResponseBody {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *CreatePublicIpAddressPoolResponseBody) SetPulbicIpAddressPoolId(v string) *CreatePublicIpAddressPoolResponseBody {
	s.PulbicIpAddressPoolId = &v
	return s
}

func (s *CreatePublicIpAddressPoolResponseBody) SetRequestId(v string) *CreatePublicIpAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublicIpAddressPoolResponseBody) SetResourceGroupId(v string) *CreatePublicIpAddressPoolResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePublicIpAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
