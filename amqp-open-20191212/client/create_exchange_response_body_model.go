// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExchangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateExchangeResponseBody
	GetRequestId() *string
}

type CreateExchangeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExchangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExchangeResponseBody) SetRequestId(v string) *CreateExchangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExchangeResponseBody) Validate() error {
	return dara.Validate(s)
}
