// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteIntentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIntentResponseBody
	GetSuccess() *bool
}

type DeleteIntentResponseBody struct {
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
	// Succes
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIntentResponseBody) SetCode(v string) *DeleteIntentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIntentResponseBody) SetHttpStatusCode(v int32) *DeleteIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIntentResponseBody) SetMessage(v string) *DeleteIntentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIntentResponseBody) SetRequestId(v string) *DeleteIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntentResponseBody) SetSuccess(v bool) *DeleteIntentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
