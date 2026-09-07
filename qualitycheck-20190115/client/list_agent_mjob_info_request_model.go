// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentMJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListAgentMJobInfoRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListAgentMJobInfoRequest
	GetJsonStr() *string
}

type ListAgentMJobInfoRequest struct {
	// The business space ID.
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following detailed information.
	//
	// example:
	//
	// {}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListAgentMJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentMJobInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAgentMJobInfoRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListAgentMJobInfoRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListAgentMJobInfoRequest) SetBaseMeAgentId(v int64) *ListAgentMJobInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListAgentMJobInfoRequest) SetJsonStr(v string) *ListAgentMJobInfoRequest {
	s.JsonStr = &v
	return s
}

func (s *ListAgentMJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
