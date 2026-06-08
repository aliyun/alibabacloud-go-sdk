// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentResponseBody
	GetCode() *string
	SetData(v *Agent) *GetAgentResponseBody
	GetData() *Agent
	SetMessage(v string) *GetAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentResponseBody
	GetSuccess() *bool
}

type GetAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *Agent `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0FFAB67-XXXXXX-CAD4D37448C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentResponseBody) GetData() *Agent {
	return s.Data
}

func (s *GetAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *Agent) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
