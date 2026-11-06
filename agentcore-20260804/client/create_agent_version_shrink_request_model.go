// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateAgentVersionShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateAgentVersionShrinkRequest
	GetClientToken() *string
}

type CreateAgentVersionShrinkRequest struct {
	// 保存版本的请求体。
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// 客户端生成的幂等令牌，当前为预留参数，后端暂不提供持久化幂等保证。
	//
	// example:
	//
	// c2d8a6f0-3b1e-4a7d-9f25-6e4b0c1a8d37
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAgentVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateAgentVersionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentVersionShrinkRequest) SetBodyShrink(v string) *CreateAgentVersionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateAgentVersionShrinkRequest) SetClientToken(v string) *CreateAgentVersionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
