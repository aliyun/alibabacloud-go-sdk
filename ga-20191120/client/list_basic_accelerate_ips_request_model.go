// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpAddress(v string) *ListBasicAccelerateIpsRequest
	GetAccelerateIpAddress() *string
	SetAccelerateIpId(v string) *ListBasicAccelerateIpsRequest
	GetAccelerateIpId() *string
	SetClientToken(v string) *ListBasicAccelerateIpsRequest
	GetClientToken() *string
	SetIpSetId(v string) *ListBasicAccelerateIpsRequest
	GetIpSetId() *string
	SetMaxResults(v int32) *ListBasicAccelerateIpsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBasicAccelerateIpsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListBasicAccelerateIpsRequest
	GetRegionId() *string
}

type ListBasicAccelerateIpsRequest struct {
	// The accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// 116.132.XX.XX
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	// The ID of the accelerated IP address of the basic GA instance.
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
	// The ID of the acceleration region.
	//
	// You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/2253380.html) operation to query the ID of the acceleration region.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query and no next queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBasicAccelerateIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsRequest) GetAccelerateIpAddress() *string {
	return s.AccelerateIpAddress
}

func (s *ListBasicAccelerateIpsRequest) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *ListBasicAccelerateIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListBasicAccelerateIpsRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *ListBasicAccelerateIpsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBasicAccelerateIpsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicAccelerateIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBasicAccelerateIpsRequest) SetAccelerateIpAddress(v string) *ListBasicAccelerateIpsRequest {
	s.AccelerateIpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetAccelerateIpId(v string) *ListBasicAccelerateIpsRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetClientToken(v string) *ListBasicAccelerateIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetIpSetId(v string) *ListBasicAccelerateIpsRequest {
	s.IpSetId = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetMaxResults(v int32) *ListBasicAccelerateIpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetNextToken(v string) *ListBasicAccelerateIpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetRegionId(v string) *ListBasicAccelerateIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) Validate() error {
	return dara.Validate(s)
}
