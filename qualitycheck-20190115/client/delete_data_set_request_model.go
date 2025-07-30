// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteDataSetRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteDataSetRequest
	GetJsonStr() *string
}

type DeleteDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"setId":"234"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteDataSetRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteDataSetRequest) SetBaseMeAgentId(v int64) *DeleteDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteDataSetRequest) SetJsonStr(v string) *DeleteDataSetRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteDataSetRequest) Validate() error {
	return dara.Validate(s)
}
