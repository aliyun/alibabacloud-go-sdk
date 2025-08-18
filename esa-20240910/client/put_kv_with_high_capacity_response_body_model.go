// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v string) *PutKvWithHighCapacityResponseBody
	GetLength() *string
	SetRequestId(v string) *PutKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetValue(v string) *PutKvWithHighCapacityResponseBody
	GetValue() *string
}

type PutKvWithHighCapacityResponseBody struct {
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

func (s PutKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityResponseBody) GetLength() *string {
	return s.Length
}

func (s *PutKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutKvWithHighCapacityResponseBody) GetValue() *string {
	return s.Value
}

func (s *PutKvWithHighCapacityResponseBody) SetLength(v string) *PutKvWithHighCapacityResponseBody {
	s.Length = &v
	return s
}

func (s *PutKvWithHighCapacityResponseBody) SetRequestId(v string) *PutKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutKvWithHighCapacityResponseBody) SetValue(v string) *PutKvWithHighCapacityResponseBody {
	s.Value = &v
	return s
}

func (s *PutKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
