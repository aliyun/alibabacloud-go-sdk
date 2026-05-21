// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTempFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTempFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTempFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTempFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTempFileResponseBody
	GetSuccess() *bool
	SetTempFileId(v string) *DeleteTempFileResponseBody
	GetTempFileId() *string
}

type DeleteTempFileResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// tempfile-05cexxxxxxxxx
	TempFileId *string `json:"TempFileId,omitempty" xml:"TempFileId,omitempty"`
}

func (s DeleteTempFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTempFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTempFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTempFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTempFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTempFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTempFileResponseBody) GetTempFileId() *string {
	return s.TempFileId
}

func (s *DeleteTempFileResponseBody) SetCode(v string) *DeleteTempFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTempFileResponseBody) SetHttpStatusCode(v int32) *DeleteTempFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTempFileResponseBody) SetMessage(v string) *DeleteTempFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTempFileResponseBody) SetRequestId(v string) *DeleteTempFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTempFileResponseBody) SetSuccess(v bool) *DeleteTempFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTempFileResponseBody) SetTempFileId(v string) *DeleteTempFileResponseBody {
	s.TempFileId = &v
	return s
}

func (s *DeleteTempFileResponseBody) Validate() error {
	return dara.Validate(s)
}
