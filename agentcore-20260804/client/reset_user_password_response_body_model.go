// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetUserPasswordResponseBody
	GetCode() *string
	SetData(v *ResetUserPasswordResponseBodyData) *ResetUserPasswordResponseBody
	GetData() *ResetUserPasswordResponseBodyData
	SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetUserPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetUserPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetUserPasswordResponseBody
	GetSuccess() *bool
}

type ResetUserPasswordResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ResetUserPasswordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetUserPasswordResponseBody) GetData() *ResetUserPasswordResponseBodyData {
	return s.Data
}

func (s *ResetUserPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetUserPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetUserPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetUserPasswordResponseBody) SetCode(v string) *ResetUserPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetData(v *ResetUserPasswordResponseBodyData) *ResetUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetUserPasswordResponseBody) SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetMessage(v string) *ResetUserPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetSuccess(v bool) *ResetUserPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ResetUserPasswordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetUserPasswordResponseBodyData struct {
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// example:
	//
	// Example@2026
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ResetUserPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyData) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *ResetUserPasswordResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ResetUserPasswordResponseBodyData) SetAgentCoreUserId(v string) *ResetUserPasswordResponseBodyData {
	s.AgentCoreUserId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetPassword(v string) *ResetUserPasswordResponseBodyData {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetWorkspaceId(v string) *ResetUserPasswordResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
