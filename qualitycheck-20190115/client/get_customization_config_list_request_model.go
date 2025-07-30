// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizationConfigListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetCustomizationConfigListRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetCustomizationConfigListRequest
	GetJsonStr() *string
}

type GetCustomizationConfigListRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetCustomizationConfigListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizationConfigListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetCustomizationConfigListRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetCustomizationConfigListRequest) SetBaseMeAgentId(v int64) *GetCustomizationConfigListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetCustomizationConfigListRequest) SetJsonStr(v string) *GetCustomizationConfigListRequest {
	s.JsonStr = &v
	return s
}

func (s *GetCustomizationConfigListRequest) Validate() error {
	return dara.Validate(s)
}
