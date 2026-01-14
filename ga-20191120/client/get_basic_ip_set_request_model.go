// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetBasicIpSetRequest
	GetClientToken() *string
	SetIpSetId(v string) *GetBasicIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *GetBasicIpSetRequest
	GetRegionId() *string
}

type GetBasicIpSetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the acceleration region.
	//
	// You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/2253380.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the basic GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *GetBasicIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicIpSetRequest) SetClientToken(v string) *GetBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicIpSetRequest) SetIpSetId(v string) *GetBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *GetBasicIpSetRequest) SetRegionId(v string) *GetBasicIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicIpSetRequest) Validate() error {
	return dara.Validate(s)
}
