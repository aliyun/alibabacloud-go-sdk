// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateTagRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateTagRequest
	GetJsonStr() *string
}

type CreateTagRequest struct {
	// The business space ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete **JSON string*	- information. For more information, see the following details.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateTagRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateTagRequest) SetBaseMeAgentId(v int64) *CreateTagRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateTagRequest) SetJsonStr(v string) *CreateTagRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
