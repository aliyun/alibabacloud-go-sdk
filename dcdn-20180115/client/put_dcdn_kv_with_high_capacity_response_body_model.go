// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *PutDcdnKvWithHighCapacityResponseBody
	GetLength() *int32
	SetRequestId(v string) *PutDcdnKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetValue(v string) *PutDcdnKvWithHighCapacityResponseBody
	GetValue() *string
}

type PutDcdnKvWithHighCapacityResponseBody struct {
	// example:
	//
	// 4
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutDcdnKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *PutDcdnKvWithHighCapacityResponseBody) GetLength() *int32 {
	return s.Length
}

func (s *PutDcdnKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDcdnKvWithHighCapacityResponseBody) GetValue() *string {
	return s.Value
}

func (s *PutDcdnKvWithHighCapacityResponseBody) SetLength(v int32) *PutDcdnKvWithHighCapacityResponseBody {
	s.Length = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponseBody) SetRequestId(v string) *PutDcdnKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponseBody) SetValue(v string) *PutDcdnKvWithHighCapacityResponseBody {
	s.Value = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
