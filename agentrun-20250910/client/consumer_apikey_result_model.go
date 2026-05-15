// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerAPIKeyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConsumerAPIKeyResult
	GetCode() *string
	SetData(v *ConsumerAPIKey) *ConsumerAPIKeyResult
	GetData() *ConsumerAPIKey
	SetRequestId(v string) *ConsumerAPIKeyResult
	GetRequestId() *string
}

type ConsumerAPIKeyResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 消费者API密钥的详细信息
	//
	// example:
	//
	// {}
	Data *ConsumerAPIKey `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ConsumerAPIKeyResult) String() string {
	return dara.Prettify(s)
}

func (s ConsumerAPIKeyResult) GoString() string {
	return s.String()
}

func (s *ConsumerAPIKeyResult) GetCode() *string {
	return s.Code
}

func (s *ConsumerAPIKeyResult) GetData() *ConsumerAPIKey {
	return s.Data
}

func (s *ConsumerAPIKeyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsumerAPIKeyResult) SetCode(v string) *ConsumerAPIKeyResult {
	s.Code = &v
	return s
}

func (s *ConsumerAPIKeyResult) SetData(v *ConsumerAPIKey) *ConsumerAPIKeyResult {
	s.Data = v
	return s
}

func (s *ConsumerAPIKeyResult) SetRequestId(v string) *ConsumerAPIKeyResult {
	s.RequestId = &v
	return s
}

func (s *ConsumerAPIKeyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
