// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetText(v string) *CreateAIStaffConversationRequest
	GetText() *string
}

type CreateAIStaffConversationRequest struct {
	// User question text. The first 100 characters are truncated and used as the session title.
	//
	// example:
	//
	// 卫浴五金产品，表面采用拉丝不锈钢材质，整体色调为现代银灰色，呈现简约且富有质感的风格。产品包括淋浴花洒、水龙头及毛巾架等配件，均具备防锈、耐腐蚀性能，适用于潮湿环境。各部件连接处设计精密，配有隐藏式螺丝结构，提升整体美观度。水龙头和花洒喷头支持多模式出水，把手操作顺滑，符合人体工学设计。所有五金件通过国家节水认证，支持冷热水接入，安装方式为壁挂式。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CreateAIStaffConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateAIStaffConversationRequest) GetText() *string {
	return s.Text
}

func (s *CreateAIStaffConversationRequest) SetText(v string) *CreateAIStaffConversationRequest {
	s.Text = &v
	return s
}

func (s *CreateAIStaffConversationRequest) Validate() error {
	return dara.Validate(s)
}
