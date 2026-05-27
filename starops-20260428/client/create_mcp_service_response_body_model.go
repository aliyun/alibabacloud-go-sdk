// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServiceName(v string) *CreateMcpServiceResponseBody
	GetMcpServiceName() *string
	SetRequestId(v string) *CreateMcpServiceResponseBody
	GetRequestId() *string
}

type CreateMcpServiceResponseBody struct {
	McpServiceName *string `json:"mcpServiceName,omitempty" xml:"mcpServiceName,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMcpServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceResponseBody) GetMcpServiceName() *string {
	return s.McpServiceName
}

func (s *CreateMcpServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpServiceResponseBody) SetMcpServiceName(v string) *CreateMcpServiceResponseBody {
	s.McpServiceName = &v
	return s
}

func (s *CreateMcpServiceResponseBody) SetRequestId(v string) *CreateMcpServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
