// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeCidrListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeNodeCidrListRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeNodeCidrListRequest
	GetOwnerId() *string
	SetRegion(v string) *DescribeNodeCidrListRequest
	GetRegion() *string
}

type DescribeNodeCidrListRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which DBS is activated. Valid values:
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-qingdao**: China (Qingdao)
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// 	- **ap-southeast-1**: Singapore (Singapore)
	//
	// 	- **cn-hangzhou-finance**: China East 1 Finance
	//
	// 	- **cn-shanghai-finance**: China East 2 Finance
	//
	// 	- **cn-shenzhen-finance**: China South 1 Finance
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeNodeCidrListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeCidrListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNodeCidrListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeNodeCidrListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeNodeCidrListRequest) SetClientToken(v string) *DescribeNodeCidrListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNodeCidrListRequest) SetOwnerId(v string) *DescribeNodeCidrListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNodeCidrListRequest) SetRegion(v string) *DescribeNodeCidrListRequest {
	s.Region = &v
	return s
}

func (s *DescribeNodeCidrListRequest) Validate() error {
	return dara.Validate(s)
}
