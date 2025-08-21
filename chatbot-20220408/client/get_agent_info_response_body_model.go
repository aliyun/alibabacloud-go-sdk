// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgentInfoResponseBodyData) *GetAgentInfoResponseBody
	GetData() *GetAgentInfoResponseBodyData
	SetMessage(v string) *GetAgentInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentInfoResponseBody
	GetSuccess() *bool
}

type GetAgentInfoResponseBody struct {
	Data *GetAgentInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponseBody) GetData() *GetAgentInfoResponseBodyData {
	return s.Data
}

func (s *GetAgentInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentInfoResponseBody) SetData(v *GetAgentInfoResponseBodyData) *GetAgentInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentInfoResponseBody) SetMessage(v string) *GetAgentInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentInfoResponseBody) SetRequestId(v string) *GetAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentInfoResponseBody) SetSuccess(v bool) *GetAgentInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAgentInfoResponseBodyData struct {
	// example:
	//
	// 4e7400028e6f4a7393ed3acf6a7b8927_p_beebot_public
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s GetAgentInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponseBodyData) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetAgentInfoResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *GetAgentInfoResponseBodyData) SetAgentKey(v string) *GetAgentInfoResponseBodyData {
	s.AgentKey = &v
	return s
}

func (s *GetAgentInfoResponseBodyData) SetAgentName(v string) *GetAgentInfoResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *GetAgentInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
