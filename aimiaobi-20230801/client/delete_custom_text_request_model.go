// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteCustomTextRequest
	GetAgentKey() *string
	SetCommodityCode(v string) *DeleteCustomTextRequest
	GetCommodityCode() *string
	SetId(v int64) *DeleteCustomTextRequest
	GetId() *int64
}

type DeleteCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 85
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCustomTextRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTextRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteCustomTextRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DeleteCustomTextRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCustomTextRequest) SetAgentKey(v string) *DeleteCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTextRequest) SetCommodityCode(v string) *DeleteCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteCustomTextRequest) SetId(v int64) *DeleteCustomTextRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomTextRequest) Validate() error {
	return dara.Validate(s)
}
