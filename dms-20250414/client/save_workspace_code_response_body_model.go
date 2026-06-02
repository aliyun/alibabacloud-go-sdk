// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkspaceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SaveWorkspaceCodeResponseBody
	GetData() *string
	SetErrorCode(v string) *SaveWorkspaceCodeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *SaveWorkspaceCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveWorkspaceCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveWorkspaceCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveWorkspaceCodeResponseBody
	GetSuccess() *bool
}

type SaveWorkspaceCodeResponseBody struct {
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Failed to publish，repo branch empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveWorkspaceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkspaceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWorkspaceCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *SaveWorkspaceCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveWorkspaceCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveWorkspaceCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveWorkspaceCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWorkspaceCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveWorkspaceCodeResponseBody) SetData(v string) *SaveWorkspaceCodeResponseBody {
	s.Data = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) SetErrorCode(v string) *SaveWorkspaceCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) SetHttpStatusCode(v int32) *SaveWorkspaceCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) SetMessage(v string) *SaveWorkspaceCodeResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) SetRequestId(v string) *SaveWorkspaceCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) SetSuccess(v bool) *SaveWorkspaceCodeResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWorkspaceCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
