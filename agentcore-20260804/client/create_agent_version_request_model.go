// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentVersionRequestBody) *CreateAgentVersionRequest
	GetBody() *CreateAgentVersionRequestBody
	SetClientToken(v string) *CreateAgentVersionRequest
	GetClientToken() *string
}

type CreateAgentVersionRequest struct {
	// 保存版本的请求体。
	Body *CreateAgentVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// 客户端生成的幂等令牌，当前为预留参数，后端暂不提供持久化幂等保证。
	//
	// example:
	//
	// c2d8a6f0-3b1e-4a7d-9f25-6e4b0c1a8d37
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAgentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionRequest) GetBody() *CreateAgentVersionRequestBody {
	return s.Body
}

func (s *CreateAgentVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentVersionRequest) SetBody(v *CreateAgentVersionRequestBody) *CreateAgentVersionRequest {
	s.Body = v
	return s
}

func (s *CreateAgentVersionRequest) SetClientToken(v string) *CreateAgentVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentVersionRequestBody struct {
	// 版本描述，最多 1024 个字符，保存后不可修改。
	//
	// example:
	//
	// 生产环境稳定版本
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 用户自定义的展示名称，去除首尾空白后为 1 至 128 个字符，可包含中文、大小写字母、空格和标点，不能包含控制字符。同一 Agent 下未删除的版本名称唯一，区分大小写，保存后不可修改。版本号 agentVersion 由系统生成，不由此字段决定。
	//
	// This parameter is required.
	//
	// example:
	//
	// 生产稳定版 V1.0
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
}

func (s CreateAgentVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionRequestBody) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentVersionRequestBody) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateAgentVersionRequestBody) SetDescription(v string) *CreateAgentVersionRequestBody {
	s.Description = &v
	return s
}

func (s *CreateAgentVersionRequestBody) SetVersionName(v string) *CreateAgentVersionRequestBody {
	s.VersionName = &v
	return s
}

func (s *CreateAgentVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
