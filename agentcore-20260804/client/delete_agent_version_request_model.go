// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteAgentVersionRequest
	GetClientToken() *string
}

type DeleteAgentVersionRequest struct {
	// 客户端生成的幂等令牌，当前为预留参数，后端暂不提供持久化幂等保证。
	//
	// example:
	//
	// c2d8a6f0-3b1e-4a7d-9f25-6e4b0c1a8d37
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteAgentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgentVersionRequest) SetClientToken(v string) *DeleteAgentVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentVersionRequest) Validate() error {
	return dara.Validate(s)
}
