// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateToolInputV2) *UpdateToolRequest
	GetBody() *UpdateToolInputV2
}

type UpdateToolRequest struct {
	// 更新工具的请求参数，只需提供需要修改的字段
	Body *UpdateToolInputV2 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateToolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolRequest) GoString() string {
	return s.String()
}

func (s *UpdateToolRequest) GetBody() *UpdateToolInputV2 {
	return s.Body
}

func (s *UpdateToolRequest) SetBody(v *UpdateToolInputV2) *UpdateToolRequest {
	s.Body = v
	return s
}

func (s *UpdateToolRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
