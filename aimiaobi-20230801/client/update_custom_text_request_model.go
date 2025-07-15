// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateCustomTextRequest
	GetAgentKey() *string
	SetCommodityCode(v string) *UpdateCustomTextRequest
	GetCommodityCode() *string
	SetContent(v string) *UpdateCustomTextRequest
	GetContent() *string
	SetId(v int64) *UpdateCustomTextRequest
	GetId() *int64
	SetTitle(v string) *UpdateCustomTextRequest
	GetTitle() *string
}

type UpdateCustomTextRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 96
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateCustomTextRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTextRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateCustomTextRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateCustomTextRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateCustomTextRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateCustomTextRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateCustomTextRequest) SetAgentKey(v string) *UpdateCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCustomTextRequest) SetCommodityCode(v string) *UpdateCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpdateCustomTextRequest) SetContent(v string) *UpdateCustomTextRequest {
	s.Content = &v
	return s
}

func (s *UpdateCustomTextRequest) SetId(v int64) *UpdateCustomTextRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomTextRequest) SetTitle(v string) *UpdateCustomTextRequest {
	s.Title = &v
	return s
}

func (s *UpdateCustomTextRequest) Validate() error {
	return dara.Validate(s)
}
