// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateTagRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateTagRequest
	GetJsonStr() *string
}

type UpdateTagRequest struct {
	// The business space ID.
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

func (s UpdateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateTagRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateTagRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateTagRequest) SetBaseMeAgentId(v int64) *UpdateTagRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateTagRequest) SetJsonStr(v string) *UpdateTagRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateTagRequest) Validate() error {
	return dara.Validate(s)
}
