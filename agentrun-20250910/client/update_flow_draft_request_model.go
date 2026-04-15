// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateFlowDraftInput) *UpdateFlowDraftRequest
	GetBody() *UpdateFlowDraftInput
}

type UpdateFlowDraftRequest struct {
	// 包含要更新的工作流草稿配置信息的请求体
	//
	// This parameter is required.
	Body *UpdateFlowDraftInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowDraftRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowDraftRequest) GetBody() *UpdateFlowDraftInput {
	return s.Body
}

func (s *UpdateFlowDraftRequest) SetBody(v *UpdateFlowDraftInput) *UpdateFlowDraftRequest {
	s.Body = v
	return s
}

func (s *UpdateFlowDraftRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
