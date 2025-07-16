// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetA1Notation(v string) *ClearDataResponseBody
	GetA1Notation() *string
	SetRequestId(v string) *ClearDataResponseBody
	GetRequestId() *string
}

type ClearDataResponseBody struct {
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

func (s ClearDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearDataResponseBody) GoString() string {
	return s.String()
}

func (s *ClearDataResponseBody) GetA1Notation() *string {
	return s.A1Notation
}

func (s *ClearDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearDataResponseBody) SetA1Notation(v string) *ClearDataResponseBody {
	s.A1Notation = &v
	return s
}

func (s *ClearDataResponseBody) SetRequestId(v string) *ClearDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearDataResponseBody) Validate() error {
	return dara.Validate(s)
}
