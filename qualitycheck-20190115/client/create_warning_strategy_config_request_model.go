// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningStrategyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateWarningStrategyConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateWarningStrategyConfigRequest
	GetJsonStr() *string
}

type CreateWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningStrategyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateWarningStrategyConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *CreateWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateWarningStrategyConfigRequest) SetJsonStr(v string) *CreateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateWarningStrategyConfigRequest) Validate() error {
	return dara.Validate(s)
}
