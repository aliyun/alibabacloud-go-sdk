// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveConfigDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SaveConfigDataSetRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SaveConfigDataSetRequest
	GetJsonStr() *string
}

type SaveConfigDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SaveConfigDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveConfigDataSetRequest) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SaveConfigDataSetRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SaveConfigDataSetRequest) SetBaseMeAgentId(v int64) *SaveConfigDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SaveConfigDataSetRequest) SetJsonStr(v string) *SaveConfigDataSetRequest {
	s.JsonStr = &v
	return s
}

func (s *SaveConfigDataSetRequest) Validate() error {
	return dara.Validate(s)
}
