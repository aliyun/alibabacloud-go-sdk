// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *GetBasicAccelerateIpRequest
	GetAccelerateIpId() *string
	SetClientToken(v string) *GetBasicAccelerateIpRequest
	GetClientToken() *string
	SetRegionId(v string) *GetBasicAccelerateIpRequest
	GetRegionId() *string
}

type GetBasicAccelerateIpRequest struct {
	// The ID of the basic GA instance that you want to query.
	//
	// You can call the [ListBasicAccelerateIps](https://help.aliyun.com/document_detail/2253393.html) operation to query the ID of the accelerated IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
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
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpRequest) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *GetBasicAccelerateIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicAccelerateIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicAccelerateIpRequest) SetAccelerateIpId(v string) *GetBasicAccelerateIpRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpRequest) SetClientToken(v string) *GetBasicAccelerateIpRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpRequest) SetRegionId(v string) *GetBasicAccelerateIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicAccelerateIpRequest) Validate() error {
	return dara.Validate(s)
}
