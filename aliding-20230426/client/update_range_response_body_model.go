// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetA1Notation(v string) *UpdateRangeResponseBody
	GetA1Notation() *string
	SetRequestId(v string) *UpdateRangeResponseBody
	GetRequestId() *string
}

type UpdateRangeResponseBody struct {
	// example:
	//
	// A1:B2
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRangeResponseBody) GetA1Notation() *string {
	return s.A1Notation
}

func (s *UpdateRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRangeResponseBody) SetA1Notation(v string) *UpdateRangeResponseBody {
	s.A1Notation = &v
	return s
}

func (s *UpdateRangeResponseBody) SetRequestId(v string) *UpdateRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRangeResponseBody) Validate() error {
	return dara.Validate(s)
}
