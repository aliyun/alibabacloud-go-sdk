// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteServiceVersionResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceVersionResponseBody
	GetRequestId() *string
}

type DeleteServiceVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceVersionResponseBody) SetCode(v string) *DeleteServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceVersionResponseBody) SetMessage(v string) *DeleteServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceVersionResponseBody) SetRequestId(v string) *DeleteServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
