// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBizEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBizEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBizEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBizEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBizEntityResponseBody
	GetSuccess() *bool
}

type DeleteBizEntityResponseBody struct {
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

func (s DeleteBizEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBizEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBizEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBizEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBizEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBizEntityResponseBody) SetCode(v string) *DeleteBizEntityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizEntityResponseBody) SetHttpStatusCode(v int32) *DeleteBizEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBizEntityResponseBody) SetMessage(v string) *DeleteBizEntityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBizEntityResponseBody) SetRequestId(v string) *DeleteBizEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBizEntityResponseBody) SetSuccess(v bool) *DeleteBizEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBizEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
