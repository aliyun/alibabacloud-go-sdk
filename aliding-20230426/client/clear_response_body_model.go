// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetA1Notation(v string) *ClearResponseBody
	GetA1Notation() *string
	SetRequestId(v string) *ClearResponseBody
	GetRequestId() *string
}

type ClearResponseBody struct {
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

func (s ClearResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearResponseBody) GoString() string {
	return s.String()
}

func (s *ClearResponseBody) GetA1Notation() *string {
	return s.A1Notation
}

func (s *ClearResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearResponseBody) SetA1Notation(v string) *ClearResponseBody {
	s.A1Notation = &v
	return s
}

func (s *ClearResponseBody) SetRequestId(v string) *ClearResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearResponseBody) Validate() error {
	return dara.Validate(s)
}
