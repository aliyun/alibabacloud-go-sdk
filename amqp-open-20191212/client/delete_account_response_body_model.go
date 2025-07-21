// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAccountResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteAccountResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAccountResponseBody
	GetSuccess() *bool
}

type DeleteAccountResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAccountResponseBody) SetCode(v int32) *DeleteAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccountResponseBody) SetData(v bool) *DeleteAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAccountResponseBody) SetMessage(v string) *DeleteAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetSuccess(v bool) *DeleteAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
