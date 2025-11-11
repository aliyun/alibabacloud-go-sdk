// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataSourceOrderConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveDataSourceOrderConfigRequest
	GetAgentKey() *string
	SetGenerateTechnology(v string) *SaveDataSourceOrderConfigRequest
	GetGenerateTechnology() *string
	SetProductCode(v string) *SaveDataSourceOrderConfigRequest
	GetProductCode() *string
	SetUserConfigDataSourceList(v []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList) *SaveDataSourceOrderConfigRequest
	GetUserConfigDataSourceList() []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList
}

type SaveDataSourceOrderConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// copilotReference
	GenerateTechnology *string `json:"GenerateTechnology,omitempty" xml:"GenerateTechnology,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	UserConfigDataSourceList []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s SaveDataSourceOrderConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveDataSourceOrderConfigRequest) GetGenerateTechnology() *string {
	return s.GenerateTechnology
}

func (s *SaveDataSourceOrderConfigRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *SaveDataSourceOrderConfigRequest) GetUserConfigDataSourceList() []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	return s.UserConfigDataSourceList
}

func (s *SaveDataSourceOrderConfigRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetGenerateTechnology(v string) *SaveDataSourceOrderConfigRequest {
	s.GenerateTechnology = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetProductCode(v string) *SaveDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetUserConfigDataSourceList(v []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList) *SaveDataSourceOrderConfigRequest {
	s.UserConfigDataSourceList = v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) Validate() error {
	if s.UserConfigDataSourceList != nil {
		for _, item := range s.UserConfigDataSourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveDataSourceOrderConfigRequestUserConfigDataSourceList struct {
	// This parameter is required.
	//
	// example:
	//
	// QuarkCommonNews
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SystemSearch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) String() string {
	return dara.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GetCode() *string {
	return s.Code
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GetEnable() *bool {
	return s.Enable
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GetName() *string {
	return s.Name
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GetNumber() *int32 {
	return s.Number
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GetType() *string {
	return s.Type
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetCode(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetEnable(v bool) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Enable = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetName(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetNumber(v int32) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetType(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Type = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) Validate() error {
	return dara.Validate(s)
}
