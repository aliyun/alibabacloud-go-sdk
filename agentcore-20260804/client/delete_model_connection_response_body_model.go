// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelConnectionResponseBody
	GetCode() *string
	SetData(v *DeleteModelConnectionResponseBodyData) *DeleteModelConnectionResponseBody
	GetData() *DeleteModelConnectionResponseBodyData
	SetHttpStatusCode(v int32) *DeleteModelConnectionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteModelConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteModelConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelConnectionResponseBody
	GetSuccess() *bool
}

type DeleteModelConnectionResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteModelConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteModelConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteModelConnectionResponseBody) GetData() *DeleteModelConnectionResponseBodyData {
	return s.Data
}

func (s *DeleteModelConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteModelConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteModelConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelConnectionResponseBody) SetCode(v string) *DeleteModelConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelConnectionResponseBody) SetData(v *DeleteModelConnectionResponseBodyData) *DeleteModelConnectionResponseBody {
	s.Data = v
	return s
}

func (s *DeleteModelConnectionResponseBody) SetHttpStatusCode(v int32) *DeleteModelConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteModelConnectionResponseBody) SetMessage(v string) *DeleteModelConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelConnectionResponseBody) SetRequestId(v string) *DeleteModelConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelConnectionResponseBody) SetSuccess(v bool) *DeleteModelConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteModelConnectionResponseBodyData struct {
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteModelConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteModelConnectionResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *DeleteModelConnectionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteModelConnectionResponseBodyData) SetConnectionId(v string) *DeleteModelConnectionResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *DeleteModelConnectionResponseBodyData) SetWorkspaceId(v string) *DeleteModelConnectionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteModelConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
