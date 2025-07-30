// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteWarningConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteWarningConfigRequest
	GetJsonStr() *string
}

type DeleteWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"configId": "31"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteWarningConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteWarningConfigRequest) SetBaseMeAgentId(v int64) *DeleteWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteWarningConfigRequest) SetJsonStr(v string) *DeleteWarningConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteWarningConfigRequest) Validate() error {
	return dara.Validate(s)
}
