// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceOrderConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetDataSourceOrderConfigRequest
	GetAgentKey() *string
	SetGenerateTechnology(v string) *GetDataSourceOrderConfigRequest
	GetGenerateTechnology() *string
	SetProductCode(v string) *GetDataSourceOrderConfigRequest
	GetProductCode() *string
}

type GetDataSourceOrderConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d9a1f6146a37446495d9985c2e7b267e_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// copilotPrecise
	GenerateTechnology *string `json:"GenerateTechnology,omitempty" xml:"GenerateTechnology,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetDataSourceOrderConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetDataSourceOrderConfigRequest) GetGenerateTechnology() *string {
	return s.GenerateTechnology
}

func (s *GetDataSourceOrderConfigRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetDataSourceOrderConfigRequest) SetAgentKey(v string) *GetDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GetDataSourceOrderConfigRequest) SetGenerateTechnology(v string) *GetDataSourceOrderConfigRequest {
	s.GenerateTechnology = &v
	return s
}

func (s *GetDataSourceOrderConfigRequest) SetProductCode(v string) *GetDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *GetDataSourceOrderConfigRequest) Validate() error {
	return dara.Validate(s)
}
