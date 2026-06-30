// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListTagRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListTagRequest
	GetJsonStr() *string
}

type ListTagRequest struct {
	// The ID of the business space.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following detailed information.
	//
	// example:
	//
	// “”
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagRequest) GoString() string {
	return s.String()
}

func (s *ListTagRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListTagRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListTagRequest) SetBaseMeAgentId(v int64) *ListTagRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListTagRequest) SetJsonStr(v string) *ListTagRequest {
	s.JsonStr = &v
	return s
}

func (s *ListTagRequest) Validate() error {
	return dara.Validate(s)
}
