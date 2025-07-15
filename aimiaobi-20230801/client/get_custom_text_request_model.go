// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetCustomTextRequest
	GetAgentKey() *string
	SetCommodityCode(v string) *GetCustomTextRequest
	GetCommodityCode() *string
	SetId(v int64) *GetCustomTextRequest
	GetId() *int64
}

type GetCustomTextRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 63
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCustomTextRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTextRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTextRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetCustomTextRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetCustomTextRequest) GetId() *int64 {
	return s.Id
}

func (s *GetCustomTextRequest) SetAgentKey(v string) *GetCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *GetCustomTextRequest) SetCommodityCode(v string) *GetCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetCustomTextRequest) SetId(v int64) *GetCustomTextRequest {
	s.Id = &v
	return s
}

func (s *GetCustomTextRequest) Validate() error {
	return dara.Validate(s)
}
