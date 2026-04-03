// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateToolInputV2) *CreateToolRequest
	GetBody() *CreateToolInputV2
}

type CreateToolRequest struct {
	// 创建工具的请求参数，包含工具的所有配置信息
	Body *CreateToolInputV2 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateToolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateToolRequest) GoString() string {
	return s.String()
}

func (s *CreateToolRequest) GetBody() *CreateToolInputV2 {
	return s.Body
}

func (s *CreateToolRequest) SetBody(v *CreateToolInputV2) *CreateToolRequest {
	s.Body = v
	return s
}

func (s *CreateToolRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
