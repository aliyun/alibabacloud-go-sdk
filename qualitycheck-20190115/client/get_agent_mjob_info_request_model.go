// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentMJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetAgentMJobInfoRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetAgentMJobInfoRequest
	GetJsonStr() *string
}

type GetAgentMJobInfoRequest struct {
	// The ID of the business workspace.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following detailed description.
	//
	// example:
	//
	// {}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAgentMJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetAgentMJobInfoRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetAgentMJobInfoRequest) SetBaseMeAgentId(v int64) *GetAgentMJobInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetAgentMJobInfoRequest) SetJsonStr(v string) *GetAgentMJobInfoRequest {
	s.JsonStr = &v
	return s
}

func (s *GetAgentMJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
