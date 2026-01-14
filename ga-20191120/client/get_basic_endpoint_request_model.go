// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetBasicEndpointRequest
	GetClientToken() *string
	SetEndpointId(v string) *GetBasicEndpointRequest
	GetEndpointId() *string
	SetRegionId(v string) *GetBasicEndpointRequest
	GetRegionId() *string
}

type GetBasicEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetBasicEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicEndpointRequest) SetClientToken(v string) *GetBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicEndpointRequest) SetEndpointId(v string) *GetBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *GetBasicEndpointRequest) SetRegionId(v string) *GetBasicEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicEndpointRequest) Validate() error {
	return dara.Validate(s)
}
