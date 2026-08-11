// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQuotaResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQuotaResponseBody
	GetRequestId() *string
}

type DeleteQuotaResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQuotaResponseBody) SetCode(v string) *DeleteQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQuotaResponseBody) SetMessage(v string) *DeleteQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQuotaResponseBody) SetRequestId(v string) *DeleteQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
