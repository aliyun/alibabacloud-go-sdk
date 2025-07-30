// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningStrategyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListWarningStrategyConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListWarningStrategyConfigRequest
	GetJsonStr() *string
}

type ListWarningStrategyConfigRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningStrategyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListWarningStrategyConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListWarningStrategyConfigRequest) SetBaseMeAgentId(v int64) *ListWarningStrategyConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListWarningStrategyConfigRequest) SetJsonStr(v string) *ListWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *ListWarningStrategyConfigRequest) Validate() error {
	return dara.Validate(s)
}
