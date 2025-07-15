// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveCustomTextRequest
	GetAgentKey() *string
	SetCommodityCode(v string) *SaveCustomTextRequest
	GetCommodityCode() *string
	SetContent(v string) *SaveCustomTextRequest
	GetContent() *string
	SetTitle(v string) *SaveCustomTextRequest
	GetTitle() *string
}

type SaveCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SaveCustomTextRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomTextRequest) GoString() string {
	return s.String()
}

func (s *SaveCustomTextRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveCustomTextRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SaveCustomTextRequest) GetContent() *string {
	return s.Content
}

func (s *SaveCustomTextRequest) GetTitle() *string {
	return s.Title
}

func (s *SaveCustomTextRequest) SetAgentKey(v string) *SaveCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveCustomTextRequest) SetCommodityCode(v string) *SaveCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveCustomTextRequest) SetContent(v string) *SaveCustomTextRequest {
	s.Content = &v
	return s
}

func (s *SaveCustomTextRequest) SetTitle(v string) *SaveCustomTextRequest {
	s.Title = &v
	return s
}

func (s *SaveCustomTextRequest) Validate() error {
	return dara.Validate(s)
}
