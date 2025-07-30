// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListDataSetRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListDataSetRequest
	GetJsonStr() *string
}

type ListDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListDataSetRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListDataSetRequest) SetBaseMeAgentId(v int64) *ListDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListDataSetRequest) SetJsonStr(v string) *ListDataSetRequest {
	s.JsonStr = &v
	return s
}

func (s *ListDataSetRequest) Validate() error {
	return dara.Validate(s)
}
