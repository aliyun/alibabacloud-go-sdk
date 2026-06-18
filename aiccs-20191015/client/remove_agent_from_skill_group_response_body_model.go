// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAgentFromSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveAgentFromSkillGroupResponseBody
	GetCode() *string
	SetData(v *RemoveAgentFromSkillGroupResponseBodyData) *RemoveAgentFromSkillGroupResponseBody
	GetData() *RemoveAgentFromSkillGroupResponseBodyData
	SetMessage(v string) *RemoveAgentFromSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveAgentFromSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveAgentFromSkillGroupResponseBody
	GetSuccess() *bool
}

type RemoveAgentFromSkillGroupResponseBody struct {
	// The status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *RemoveAgentFromSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveAgentFromSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAgentFromSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAgentFromSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveAgentFromSkillGroupResponseBody) GetData() *RemoveAgentFromSkillGroupResponseBodyData {
	return s.Data
}

func (s *RemoveAgentFromSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveAgentFromSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAgentFromSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveAgentFromSkillGroupResponseBody) SetCode(v string) *RemoveAgentFromSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBody) SetData(v *RemoveAgentFromSkillGroupResponseBodyData) *RemoveAgentFromSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBody) SetMessage(v string) *RemoveAgentFromSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBody) SetRequestId(v string) *RemoveAgentFromSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBody) SetSuccess(v bool) *RemoveAgentFromSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveAgentFromSkillGroupResponseBodyData struct {
	// The number of agents successfully removed.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s RemoveAgentFromSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveAgentFromSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveAgentFromSkillGroupResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *RemoveAgentFromSkillGroupResponseBodyData) SetSuccessCount(v int32) *RemoveAgentFromSkillGroupResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *RemoveAgentFromSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
