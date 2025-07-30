// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrecisionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListPrecisionTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListPrecisionTaskRequest
	GetJsonStr() *string
}

type ListPrecisionTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"pageSize":10,"pageNumber":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListPrecisionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListPrecisionTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListPrecisionTaskRequest) SetBaseMeAgentId(v int64) *ListPrecisionTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListPrecisionTaskRequest) SetJsonStr(v string) *ListPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *ListPrecisionTaskRequest) Validate() error {
	return dara.Validate(s)
}
