// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLogStoreFromEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DetachLogStoreFromEndpointGroupRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DetachLogStoreFromEndpointGroupRequest
	GetClientToken() *string
	SetEndpointGroupIds(v []*string) *DetachLogStoreFromEndpointGroupRequest
	GetEndpointGroupIds() []*string
	SetListenerId(v string) *DetachLogStoreFromEndpointGroupRequest
	GetListenerId() *string
	SetRegionId(v string) *DetachLogStoreFromEndpointGroupRequest
	GetRegionId() *string
}

type DetachLogStoreFromEndpointGroupRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of endpoint groups.
	//
	// This parameter is required.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DetachLogStoreFromEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachLogStoreFromEndpointGroupRequest) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *DetachLogStoreFromEndpointGroupRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DetachLogStoreFromEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetAcceleratorId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetClientToken(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetEndpointGroupIds(v []*string) *DetachLogStoreFromEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetListenerId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetRegionId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
