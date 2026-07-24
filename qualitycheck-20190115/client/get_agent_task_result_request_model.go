// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v string) *GetAgentTaskResultRequest
	GetBaseMeAgentId() *string
	SetJsonStr(v string) *GetAgentTaskResultRequest
	GetJsonStr() *string
}

type GetAgentTaskResultRequest struct {
	// The ID of the business workspace.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *string `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// example:
	//
	// {\\"vid\\":\\"sip-11-1766561862.293393\\",\\"taskId\\":\\"20251224-D3B32484-2D53-5B53-A618-483A7941029E\\"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAgentTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultRequest) GetBaseMeAgentId() *string {
	return s.BaseMeAgentId
}

func (s *GetAgentTaskResultRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetAgentTaskResultRequest) SetBaseMeAgentId(v string) *GetAgentTaskResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetAgentTaskResultRequest) SetJsonStr(v string) *GetAgentTaskResultRequest {
	s.JsonStr = &v
	return s
}

func (s *GetAgentTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
