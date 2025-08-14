// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetNacosMcpServerRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *GetNacosMcpServerRequest
	GetInstanceId() *string
	SetMcpServerId(v string) *GetNacosMcpServerRequest
	GetMcpServerId() *string
	SetMcpServerVersion(v string) *GetNacosMcpServerRequest
	GetMcpServerVersion() *string
	SetNamespaceId(v string) *GetNacosMcpServerRequest
	GetNamespaceId() *string
}

type GetNacosMcpServerRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// mse_prepaid_public_cn-tl327w****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5e3c5211-d365-4013-8540-c4106ec2a5dc
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	// example:
	//
	// 1.0.0
	McpServerVersion *string `json:"McpServerVersion,omitempty" xml:"McpServerVersion,omitempty"`
	// example:
	//
	// 5e3ee449-b5c0-4aee-b857-32c0acbebf26
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetNacosMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerRequest) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetNacosMcpServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNacosMcpServerRequest) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetNacosMcpServerRequest) GetMcpServerVersion() *string {
	return s.McpServerVersion
}

func (s *GetNacosMcpServerRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNacosMcpServerRequest) SetAcceptLanguage(v string) *GetNacosMcpServerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetNacosMcpServerRequest) SetInstanceId(v string) *GetNacosMcpServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNacosMcpServerRequest) SetMcpServerId(v string) *GetNacosMcpServerRequest {
	s.McpServerId = &v
	return s
}

func (s *GetNacosMcpServerRequest) SetMcpServerVersion(v string) *GetNacosMcpServerRequest {
	s.McpServerVersion = &v
	return s
}

func (s *GetNacosMcpServerRequest) SetNamespaceId(v string) *GetNacosMcpServerRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNacosMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
