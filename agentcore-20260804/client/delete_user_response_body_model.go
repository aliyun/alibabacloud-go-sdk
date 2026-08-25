// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUserResponseBody
	GetCode() *string
	SetData(v *DeleteUserResponseBodyData) *DeleteUserResponseBody
	GetData() *DeleteUserResponseBodyData
	SetHttpStatusCode(v int32) *DeleteUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserResponseBody
	GetSuccess() *bool
}

type DeleteUserResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s DeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUserResponseBody) GetData() *DeleteUserResponseBodyData {
	return s.Data
}

func (s *DeleteUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetData(v *DeleteUserResponseBodyData) *DeleteUserResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUserResponseBody) SetHttpStatusCode(v int32) *DeleteUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteUserResponseBodyData struct {
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBodyData) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *DeleteUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteUserResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteUserResponseBodyData) SetAgentCoreUserId(v string) *DeleteUserResponseBodyData {
	s.AgentCoreUserId = &v
	return s
}

func (s *DeleteUserResponseBodyData) SetName(v string) *DeleteUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteUserResponseBodyData) SetWorkspaceId(v string) *DeleteUserResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
