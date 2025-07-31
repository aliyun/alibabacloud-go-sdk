// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDirectoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDirectoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDirectoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDirectoryResponseBody
	GetSuccess() *bool
}

type DeleteDirectoryResponseBody struct {
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

func (s DeleteDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDirectoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDirectoryResponseBody) SetCode(v string) *DeleteDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetHttpStatusCode(v int32) *DeleteDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetMessage(v string) *DeleteDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetRequestId(v string) *DeleteDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetSuccess(v bool) *DeleteDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
