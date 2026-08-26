// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkspaceResponseBody
	GetCode() *string
	SetData(v *DeleteWorkspaceResponseBodyData) *DeleteWorkspaceResponseBody
	GetData() *DeleteWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *DeleteWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkspaceResponseBody
	GetSuccess() *bool
}

type DeleteWorkspaceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result of the delete request.
	Data *DeleteWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkspaceResponseBody) GetData() *DeleteWorkspaceResponseBodyData {
	return s.Data
}

func (s *DeleteWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkspaceResponseBody) SetCode(v string) *DeleteWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetData(v *DeleteWorkspaceResponseBodyData) *DeleteWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetHttpStatusCode(v int32) *DeleteWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetMessage(v string) *DeleteWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetSuccess(v bool) *DeleteWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteWorkspaceResponseBodyData struct {
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceResponseBodyData) SetWorkspaceId(v string) *DeleteWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
