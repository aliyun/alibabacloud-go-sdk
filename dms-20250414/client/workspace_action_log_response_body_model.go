// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *WorkspaceActionLogResponseBody
	GetData() *string
	SetErrorCode(v string) *WorkspaceActionLogResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *WorkspaceActionLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *WorkspaceActionLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *WorkspaceActionLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WorkspaceActionLogResponseBody
	GetSuccess() *bool
}

type WorkspaceActionLogResponseBody struct {
	// example:
	//
	// log info
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WorkspaceActionLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionLogResponseBody) GoString() string {
	return s.String()
}

func (s *WorkspaceActionLogResponseBody) GetData() *string {
	return s.Data
}

func (s *WorkspaceActionLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WorkspaceActionLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WorkspaceActionLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WorkspaceActionLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WorkspaceActionLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WorkspaceActionLogResponseBody) SetData(v string) *WorkspaceActionLogResponseBody {
	s.Data = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) SetErrorCode(v string) *WorkspaceActionLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) SetHttpStatusCode(v int32) *WorkspaceActionLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) SetMessage(v string) *WorkspaceActionLogResponseBody {
	s.Message = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) SetRequestId(v string) *WorkspaceActionLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) SetSuccess(v bool) *WorkspaceActionLogResponseBody {
	s.Success = &v
	return s
}

func (s *WorkspaceActionLogResponseBody) Validate() error {
	return dara.Validate(s)
}
