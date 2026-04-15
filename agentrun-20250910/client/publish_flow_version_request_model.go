// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PublishFlowVersionInput) *PublishFlowVersionRequest
	GetBody() *PublishFlowVersionInput
}

type PublishFlowVersionRequest struct {
	// 包含版本描述等信息的请求体
	Body *PublishFlowVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishFlowVersionRequest) GetBody() *PublishFlowVersionInput {
	return s.Body
}

func (s *PublishFlowVersionRequest) SetBody(v *PublishFlowVersionInput) *PublishFlowVersionRequest {
	s.Body = v
	return s
}

func (s *PublishFlowVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
