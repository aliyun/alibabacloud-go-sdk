// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v string) *PutKvResponseBody
	GetLength() *string
	SetRequestId(v string) *PutKvResponseBody
	GetRequestId() *string
	SetValue(v string) *PutKvResponseBody
	GetValue() *string
}

type PutKvResponseBody struct {
	// The length of the value in the key-value pair.
	//
	// example:
	//
	// 4
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the key. If the content has more than 256 characters in length, the system displays the first 100 and the last 100 characters, and omits the middle part.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutKvResponseBody) GoString() string {
	return s.String()
}

func (s *PutKvResponseBody) GetLength() *string {
	return s.Length
}

func (s *PutKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutKvResponseBody) GetValue() *string {
	return s.Value
}

func (s *PutKvResponseBody) SetLength(v string) *PutKvResponseBody {
	s.Length = &v
	return s
}

func (s *PutKvResponseBody) SetRequestId(v string) *PutKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutKvResponseBody) SetValue(v string) *PutKvResponseBody {
	s.Value = &v
	return s
}

func (s *PutKvResponseBody) Validate() error {
	return dara.Validate(s)
}
