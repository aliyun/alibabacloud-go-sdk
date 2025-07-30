// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrecisionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeletePrecisionTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeletePrecisionTaskRequest
	GetJsonStr() *string
}

type DeletePrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId": "7C1DEF5F-2C18-4D36-99C6*******"}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeletePrecisionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeletePrecisionTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeletePrecisionTaskRequest) SetBaseMeAgentId(v int64) *DeletePrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeletePrecisionTaskRequest) SetJsonStr(v string) *DeletePrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *DeletePrecisionTaskRequest) Validate() error {
	return dara.Validate(s)
}
