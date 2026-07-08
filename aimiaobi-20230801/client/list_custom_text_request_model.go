// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListCustomTextRequest
	GetAgentKey() *string
	SetCommodityCode(v string) *ListCustomTextRequest
	GetCommodityCode() *string
}

type ListCustomTextRequest struct {
	// Unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Commodity code.
	//
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListCustomTextRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTextRequest) GoString() string {
	return s.String()
}

func (s *ListCustomTextRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListCustomTextRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListCustomTextRequest) SetAgentKey(v string) *ListCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomTextRequest) SetCommodityCode(v string) *ListCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *ListCustomTextRequest) Validate() error {
	return dara.Validate(s)
}
