// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningStrategyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetWarningStrategyConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetWarningStrategyConfigRequest
	GetJsonStr() *string
}

type GetWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetWarningStrategyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetWarningStrategyConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *GetWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetWarningStrategyConfigRequest) SetJsonStr(v string) *GetWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *GetWarningStrategyConfigRequest) Validate() error {
	return dara.Validate(s)
}
