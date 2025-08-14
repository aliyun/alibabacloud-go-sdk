// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNacosMcpServerRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *DeleteNacosMcpServerRequest
	GetInstanceId() *string
	SetMcpServerId(v string) *DeleteNacosMcpServerRequest
	GetMcpServerId() *string
	SetNamespaceId(v string) *DeleteNacosMcpServerRequest
	GetNamespaceId() *string
}

type DeleteNacosMcpServerRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// mse_prepaid_public_cn-u0t2xzvxa06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5e3c5211-d365-4013-8540-c4106ec2a5dc
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	// example:
	//
	// fc0f6e40-****-946b-45e3af313707
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNacosMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosMcpServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosMcpServerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNacosMcpServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNacosMcpServerRequest) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *DeleteNacosMcpServerRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNacosMcpServerRequest) SetAcceptLanguage(v string) *DeleteNacosMcpServerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNacosMcpServerRequest) SetInstanceId(v string) *DeleteNacosMcpServerRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosMcpServerRequest) SetMcpServerId(v string) *DeleteNacosMcpServerRequest {
	s.McpServerId = &v
	return s
}

func (s *DeleteNacosMcpServerRequest) SetNamespaceId(v string) *DeleteNacosMcpServerRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
