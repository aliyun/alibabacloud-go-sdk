// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServiceName(v string) *UpdateMcpServiceResponseBody
	GetMcpServiceName() *string
	SetRequestId(v string) *UpdateMcpServiceResponseBody
	GetRequestId() *string
}

type UpdateMcpServiceResponseBody struct {
	// example:
	//
	// log-query
	McpServiceName *string `json:"mcpServiceName,omitempty" xml:"mcpServiceName,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMcpServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceResponseBody) GetMcpServiceName() *string {
	return s.McpServiceName
}

func (s *UpdateMcpServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpServiceResponseBody) SetMcpServiceName(v string) *UpdateMcpServiceResponseBody {
	s.McpServiceName = &v
	return s
}

func (s *UpdateMcpServiceResponseBody) SetRequestId(v string) *UpdateMcpServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
