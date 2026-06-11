// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetWorkspaceCodeResponseBody
	GetData() *string
	SetErrorCode(v string) *GetWorkspaceCodeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetWorkspaceCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspaceCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceCodeResponseBody
	GetSuccess() *bool
}

type GetWorkspaceCodeResponseBody struct {
	// The file content.
	//
	// example:
	//
	// log info
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkspaceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *GetWorkspaceCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkspaceCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceCodeResponseBody) SetData(v string) *GetWorkspaceCodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) SetErrorCode(v string) *GetWorkspaceCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) SetHttpStatusCode(v int32) *GetWorkspaceCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) SetMessage(v string) *GetWorkspaceCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) SetRequestId(v string) *GetWorkspaceCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) SetSuccess(v bool) *GetWorkspaceCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
