// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningStrategyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteWarningStrategyConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteWarningStrategyConfigRequest
	GetJsonStr() *string
}

type DeleteWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningStrategyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteWarningStrategyConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *DeleteWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteWarningStrategyConfigRequest) SetJsonStr(v string) *DeleteWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteWarningStrategyConfigRequest) Validate() error {
	return dara.Validate(s)
}
