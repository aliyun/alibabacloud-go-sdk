// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetTagRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetTagRequest
	GetJsonStr() *string
}

type GetTagRequest struct {
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

func (s GetTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTagRequest) GoString() string {
	return s.String()
}

func (s *GetTagRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetTagRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetTagRequest) SetBaseMeAgentId(v int64) *GetTagRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetTagRequest) SetJsonStr(v string) *GetTagRequest {
	s.JsonStr = &v
	return s
}

func (s *GetTagRequest) Validate() error {
	return dara.Validate(s)
}
