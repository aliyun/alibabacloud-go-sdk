// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAPIKeyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerAPIKeyResult
	GetCode() *string
	SetData(v *CreateConsumerAPIKeyOutput) *CreateConsumerAPIKeyResult
	GetData() *CreateConsumerAPIKeyOutput
	SetRequestId(v string) *CreateConsumerAPIKeyResult
	GetRequestId() *string
}

type CreateConsumerAPIKeyResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建的消费者API密钥详细信息（包含完整密钥）
	//
	// example:
	//
	// {}
	Data *CreateConsumerAPIKeyOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerAPIKeyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAPIKeyResult) GoString() string {
	return s.String()
}

func (s *CreateConsumerAPIKeyResult) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerAPIKeyResult) GetData() *CreateConsumerAPIKeyOutput {
	return s.Data
}

func (s *CreateConsumerAPIKeyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerAPIKeyResult) SetCode(v string) *CreateConsumerAPIKeyResult {
	s.Code = &v
	return s
}

func (s *CreateConsumerAPIKeyResult) SetData(v *CreateConsumerAPIKeyOutput) *CreateConsumerAPIKeyResult {
	s.Data = v
	return s
}

func (s *CreateConsumerAPIKeyResult) SetRequestId(v string) *CreateConsumerAPIKeyResult {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerAPIKeyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
