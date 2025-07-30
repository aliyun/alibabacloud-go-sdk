// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningStrategyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateWarningStrategyConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateWarningStrategyConfigRequest
	GetJsonStr() *string
}

type UpdateWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningStrategyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateWarningStrategyConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *UpdateWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateWarningStrategyConfigRequest) SetJsonStr(v string) *UpdateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateWarningStrategyConfigRequest) Validate() error {
	return dara.Validate(s)
}
