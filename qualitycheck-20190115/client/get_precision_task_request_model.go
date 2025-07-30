// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrecisionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetPrecisionTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetPrecisionTaskRequest
	GetJsonStr() *string
}

type GetPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId":"593A04C0-E6E9-4CDC-8C9****"}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetPrecisionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetPrecisionTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetPrecisionTaskRequest) SetBaseMeAgentId(v int64) *GetPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetPrecisionTaskRequest) SetJsonStr(v string) *GetPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *GetPrecisionTaskRequest) Validate() error {
	return dara.Validate(s)
}
