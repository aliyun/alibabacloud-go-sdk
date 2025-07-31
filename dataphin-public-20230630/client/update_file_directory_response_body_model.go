// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFileDirectoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateFileDirectoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateFileDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFileDirectoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFileDirectoryResponseBody
	GetSuccess() *bool
}

type UpdateFileDirectoryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFileDirectoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFileDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFileDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFileDirectoryResponseBody) SetCode(v string) *UpdateFileDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetHttpStatusCode(v int32) *UpdateFileDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetMessage(v string) *UpdateFileDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetRequestId(v string) *UpdateFileDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetSuccess(v bool) *UpdateFileDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
