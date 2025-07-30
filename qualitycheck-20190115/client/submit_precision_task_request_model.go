// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPrecisionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SubmitPrecisionTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SubmitPrecisionTaskRequest
	GetJsonStr() *string
}

type SubmitPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"name":"test","dataSetId":1865}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitPrecisionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SubmitPrecisionTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SubmitPrecisionTaskRequest) SetBaseMeAgentId(v int64) *SubmitPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitPrecisionTaskRequest) SetJsonStr(v string) *SubmitPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *SubmitPrecisionTaskRequest) Validate() error {
	return dara.Validate(s)
}
