// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdatePortfolioRequest
	GetDescription() *string
	SetPortfolioId(v string) *UpdatePortfolioRequest
	GetPortfolioId() *string
	SetPortfolioName(v string) *UpdatePortfolioRequest
	GetPortfolioName() *string
	SetProviderName(v string) *UpdatePortfolioRequest
	GetProviderName() *string
}

type UpdatePortfolioRequest struct {
	// 产品组合描述
	//
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// This parameter is required.
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// 产品组合提供方
	//
	// This parameter is required.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s UpdatePortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortfolioRequest) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *UpdatePortfolioRequest) GetPortfolioName() *string {
	return s.PortfolioName
}

func (s *UpdatePortfolioRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *UpdatePortfolioRequest) SetDescription(v string) *UpdatePortfolioRequest {
	s.Description = &v
	return s
}

func (s *UpdatePortfolioRequest) SetPortfolioId(v string) *UpdatePortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *UpdatePortfolioRequest) SetPortfolioName(v string) *UpdatePortfolioRequest {
	s.PortfolioName = &v
	return s
}

func (s *UpdatePortfolioRequest) SetProviderName(v string) *UpdatePortfolioRequest {
	s.ProviderName = &v
	return s
}

func (s *UpdatePortfolioRequest) Validate() error {
	return dara.Validate(s)
}
