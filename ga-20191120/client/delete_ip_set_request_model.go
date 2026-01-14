// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteIpSetRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteIpSetRequest
	GetClientToken() *string
	SetIpSetId(v string) *DeleteIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *DeleteIpSetRequest
	GetRegionId() *string
}

type DeleteIpSetRequest struct {
	// The ID of the GA instance for which you want to delete an acceleration region.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmqy****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// DD61839A-5CC5-404B-8C6E-56066F0C432D
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the acceleration region that you want to delete.
	//
	// You can call the [ListIpSets](https://help.aliyun.com/document_detail/2253273.html) operation to query the IDs of acceleration regions of a specified GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *DeleteIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpSetRequest) SetAcceleratorId(v string) *DeleteIpSetRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteIpSetRequest) SetClientToken(v string) *DeleteIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpSetRequest) SetIpSetId(v string) *DeleteIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DeleteIpSetRequest) SetRegionId(v string) *DeleteIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpSetRequest) Validate() error {
	return dara.Validate(s)
}
