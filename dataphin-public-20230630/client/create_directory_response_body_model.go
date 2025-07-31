// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDirectoryResponseBody
	GetCode() *string
	SetFileId(v int64) *CreateDirectoryResponseBody
	GetFileId() *int64
	SetHttpStatusCode(v int32) *CreateDirectoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDirectoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDirectoryResponseBody
	GetSuccess() *bool
}

type CreateDirectoryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1311113211
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDirectoryResponseBody) GetFileId() *int64 {
	return s.FileId
}

func (s *CreateDirectoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDirectoryResponseBody) SetCode(v string) *CreateDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetFileId(v int64) *CreateDirectoryResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetHttpStatusCode(v int32) *CreateDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetMessage(v string) *CreateDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetSuccess(v bool) *CreateDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
