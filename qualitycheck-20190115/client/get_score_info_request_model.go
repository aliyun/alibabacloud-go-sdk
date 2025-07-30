// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScoreInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetScoreInfoRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetScoreInfoRequest
	GetJsonStr() *string
}

type GetScoreInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetScoreInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScoreInfoRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetScoreInfoRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetScoreInfoRequest) SetBaseMeAgentId(v int64) *GetScoreInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetScoreInfoRequest) SetJsonStr(v string) *GetScoreInfoRequest {
	s.JsonStr = &v
	return s
}

func (s *GetScoreInfoRequest) Validate() error {
	return dara.Validate(s)
}
