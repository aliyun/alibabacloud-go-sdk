// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardProtocol(v string) *DescribeLayer4RulesRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *DescribeLayer4RulesRequest
	GetFrontendPort() *int32
	SetInstanceId(v string) *DescribeLayer4RulesRequest
	GetInstanceId() *string
	SetOffset(v int32) *DescribeLayer4RulesRequest
	GetOffset() *int32
	SetPageSize(v string) *DescribeLayer4RulesRequest
	GetPageSize() *string
	SetSourceIp(v string) *DescribeLayer4RulesRequest
	GetSourceIp() *string
}

type DescribeLayer4RulesRequest struct {
	// The type of forwarding protocol. Values:
	//
	// - **tcp**: Indicates TCP protocol.
	//
	// - **udp**: Indicates UDP protocol.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the DDoS protection instance to be queried.
	//
	// > You can call [DescribeInstances](https://help.aliyun.com/document_detail/91478.html) to query all DDoS protection instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-zvp2ay9b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// In paginated queries, specifies which page of data to return. The minimum value is **1**, indicating the first page of data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// In paginated queries, specifies the number of results per page. The maximum value is **50**, indicating that each page can contain up to 50 results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request. You do not need to fill this in; it is automatically obtained by the system.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer4RulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *DescribeLayer4RulesRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RulesRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeLayer4RulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLayer4RulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLayer4RulesRequest) SetForwardProtocol(v string) *DescribeLayer4RulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetFrontendPort(v int32) *DescribeLayer4RulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetInstanceId(v string) *DescribeLayer4RulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetOffset(v int32) *DescribeLayer4RulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetPageSize(v string) *DescribeLayer4RulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetSourceIp(v string) *DescribeLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RulesRequest) Validate() error {
	return dara.Validate(s)
}
