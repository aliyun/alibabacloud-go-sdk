// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetAgentRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetAgentRequest
	GetJsonStr() *string
}

type GetAgentRequest struct {
	// baseMeAgentId
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetAgentRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetAgentRequest) SetBaseMeAgentId(v int64) *GetAgentRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetAgentRequest) SetJsonStr(v string) *GetAgentRequest {
	s.JsonStr = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
