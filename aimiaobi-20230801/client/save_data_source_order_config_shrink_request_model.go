// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataSourceOrderConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveDataSourceOrderConfigShrinkRequest
	GetAgentKey() *string
	SetGenerateTechnology(v string) *SaveDataSourceOrderConfigShrinkRequest
	GetGenerateTechnology() *string
	SetProductCode(v string) *SaveDataSourceOrderConfigShrinkRequest
	GetProductCode() *string
	SetUserConfigDataSourceListShrink(v string) *SaveDataSourceOrderConfigShrinkRequest
	GetUserConfigDataSourceListShrink() *string
}

type SaveDataSourceOrderConfigShrinkRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The generation technology for Miaosou. This parameter is valid only when \\`ProductCode\\` is set to \\`miaosou\\`.
	//
	// Valid values:
	//
	// - copilotPrecise
	//
	// - copilotReference
	//
	// example:
	//
	// copilotReference
	GenerateTechnology *string `json:"GenerateTechnology,omitempty" xml:"GenerateTechnology,omitempty"`
	// The product type. Valid values:
	//
	// - miaobi
	//
	// - miaosou
	//
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The data source configuration.
	//
	// This parameter is required.
	UserConfigDataSourceListShrink *string `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty"`
}

func (s SaveDataSourceOrderConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveDataSourceOrderConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveDataSourceOrderConfigShrinkRequest) GetGenerateTechnology() *string {
	return s.GenerateTechnology
}

func (s *SaveDataSourceOrderConfigShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *SaveDataSourceOrderConfigShrinkRequest) GetUserConfigDataSourceListShrink() *string {
	return s.UserConfigDataSourceListShrink
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetGenerateTechnology(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.GenerateTechnology = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetProductCode(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetUserConfigDataSourceListShrink(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.UserConfigDataSourceListShrink = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
