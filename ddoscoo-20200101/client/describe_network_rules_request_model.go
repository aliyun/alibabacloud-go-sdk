// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardProtocol(v string) *DescribeNetworkRulesRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *DescribeNetworkRulesRequest
	GetFrontendPort() *int32
	SetInstanceId(v string) *DescribeNetworkRulesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeNetworkRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkRulesRequest
	GetPageSize() *int32
}

type DescribeNetworkRulesRequest struct {
	// The forwarding protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 80
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworkRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *DescribeNetworkRulesRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeNetworkRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkRulesRequest) SetForwardProtocol(v string) *DescribeNetworkRulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetFrontendPort(v int32) *DescribeNetworkRulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetInstanceId(v string) *DescribeNetworkRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetPageNumber(v int32) *DescribeNetworkRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetPageSize(v int32) *DescribeNetworkRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkRulesRequest) Validate() error {
	return dara.Validate(s)
}
