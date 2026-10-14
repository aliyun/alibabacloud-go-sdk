// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrontendPort(v int32) *DescribePortRequest
	GetFrontendPort() *int32
	SetFrontendProtocol(v string) *DescribePortRequest
	GetFrontendProtocol() *string
	SetInstanceId(v string) *DescribePortRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribePortRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePortRequest
	GetPageSize() *int32
}

type DescribePortRequest struct {
	// The forward port to query. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 55
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The Forwarding Protocol type to query. Valid values:
	//
	// - **tcp**: TCP protocol.
	//
	// - **udp**: UDP protocol.
	//
	// example:
	//
	// tcp
	FrontendProtocol *string `json:"FrontendProtocol,omitempty" xml:"FrontendProtocol,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to query.
	//
	// > You can invoke [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-7e225i41****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return in a paged query. For example, set this parameter to **1*	- to return the first page of results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on each page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePortRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribePortRequest) GetFrontendProtocol() *string {
	return s.FrontendProtocol
}

func (s *DescribePortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePortRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePortRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePortRequest) SetFrontendPort(v int32) *DescribePortRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribePortRequest) SetFrontendProtocol(v string) *DescribePortRequest {
	s.FrontendProtocol = &v
	return s
}

func (s *DescribePortRequest) SetInstanceId(v string) *DescribePortRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePortRequest) SetPageNumber(v int32) *DescribePortRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePortRequest) SetPageSize(v int32) *DescribePortRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePortRequest) Validate() error {
	return dara.Validate(s)
}
