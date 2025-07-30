// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteCustomizationConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteCustomizationConfigRequest
	GetJsonStr() *string
}

type DeleteCustomizationConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"modelId":"2412"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteCustomizationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizationConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteCustomizationConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteCustomizationConfigRequest) SetBaseMeAgentId(v int64) *DeleteCustomizationConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteCustomizationConfigRequest) SetJsonStr(v string) *DeleteCustomizationConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteCustomizationConfigRequest) Validate() error {
	return dara.Validate(s)
}
