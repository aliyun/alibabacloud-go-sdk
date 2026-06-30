// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteTagRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteTagRequest
	GetJsonStr() *string
}

type DeleteTagRequest struct {
	// The ID of the business space.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A complete JSON string. For more information, see the following detailed information.
	//
	// example:
	//
	// “”
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteTagRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteTagRequest) SetBaseMeAgentId(v int64) *DeleteTagRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteTagRequest) SetJsonStr(v string) *DeleteTagRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
