// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpIdleCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetBasicAccelerateIpIdleCountRequest
	GetClientToken() *string
	SetIpSetId(v string) *GetBasicAccelerateIpIdleCountRequest
	GetIpSetId() *string
	SetRegionId(v string) *GetBasicAccelerateIpIdleCountRequest
	GetRegionId() *string
}

type GetBasicAccelerateIpIdleCountRequest struct {
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
	// The ID of the basic GA instance that you want to query.
	//
	// You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/2253380.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpIdleCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicAccelerateIpIdleCountRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *GetBasicAccelerateIpIdleCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetClientToken(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetIpSetId(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.IpSetId = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetRegionId(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountRequest) Validate() error {
	return dara.Validate(s)
}
