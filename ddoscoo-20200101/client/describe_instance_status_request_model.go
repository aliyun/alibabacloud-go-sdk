// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceStatusRequest
	GetInstanceId() *string
	SetProductType(v int32) *DescribeInstanceStatusRequest
	GetProductType() *int32
}

type DescribeInstanceStatusRequest struct {
	// The ID of the Anti-DDoS Proxy instance to query.
	//
	// >  You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all Anti-DDoS Proxy (Chinese Mainland) or Anti-DDoS Proxy (Outside Chinese Mainland) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-6ja1y6p5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the Anti-DDoS Proxy instance to query. Valid values:
	//
	// 	- **1**: an Anti-DDoS Proxy (Chinese Mainland) instance
	//
	// 	- **2**: an Anti-DDoS Proxy (Outside Chinese Mainland) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int32 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceStatusRequest) GetProductType() *int32 {
	return s.ProductType
}

func (s *DescribeInstanceStatusRequest) SetInstanceId(v string) *DescribeInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetProductType(v int32) *DescribeInstanceStatusRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
