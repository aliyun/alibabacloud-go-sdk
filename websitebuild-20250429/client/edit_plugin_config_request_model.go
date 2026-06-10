// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPluginConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizId(v string) *EditPluginConfigRequest
  GetBizId() *string 
  SetPluginConfig(v string) *EditPluginConfigRequest
  GetPluginConfig() *string 
  SetPluginDesc(v string) *EditPluginConfigRequest
  GetPluginDesc() *string 
  SetPluginId(v string) *EditPluginConfigRequest
  GetPluginId() *string 
  SetPluginName(v string) *EditPluginConfigRequest
  GetPluginName() *string 
}

type EditPluginConfigRequest struct {
  // Business ID
  // 
  // example:
  // 
  // WD20250703155602000001
  BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
  // Plugin configuration
  // 
  // example:
  // 
  // {}
  PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
  // Plugin description
  // 
  // example:
  // 
  // 根据关键字查询相关的Github Repositories
  PluginDesc *string `json:"PluginDesc,omitempty" xml:"PluginDesc,omitempty"`
  // Plugin ID
  // 
  // example:
  // 
  // aliplayer-react
  PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
  // Plugin name
  // 
  // example:
  // 
  // alisecguard
  PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
}

func (s EditPluginConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s EditPluginConfigRequest) GoString() string {
  return s.String()
}

func (s *EditPluginConfigRequest) GetBizId() *string  {
  return s.BizId
}

func (s *EditPluginConfigRequest) GetPluginConfig() *string  {
  return s.PluginConfig
}

func (s *EditPluginConfigRequest) GetPluginDesc() *string  {
  return s.PluginDesc
}

func (s *EditPluginConfigRequest) GetPluginId() *string  {
  return s.PluginId
}

func (s *EditPluginConfigRequest) GetPluginName() *string  {
  return s.PluginName
}

func (s *EditPluginConfigRequest) SetBizId(v string) *EditPluginConfigRequest {
  s.BizId = &v
  return s
}

func (s *EditPluginConfigRequest) SetPluginConfig(v string) *EditPluginConfigRequest {
  s.PluginConfig = &v
  return s
}

func (s *EditPluginConfigRequest) SetPluginDesc(v string) *EditPluginConfigRequest {
  s.PluginDesc = &v
  return s
}

func (s *EditPluginConfigRequest) SetPluginId(v string) *EditPluginConfigRequest {
  s.PluginId = &v
  return s
}

func (s *EditPluginConfigRequest) SetPluginName(v string) *EditPluginConfigRequest {
  s.PluginName = &v
  return s
}

func (s *EditPluginConfigRequest) Validate() error {
  return dara.Validate(s)
}

