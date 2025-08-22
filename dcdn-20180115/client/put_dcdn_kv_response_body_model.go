// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *PutDcdnKvResponseBody
	GetLength() *int32
	SetRequestId(v string) *PutDcdnKvResponseBody
	GetRequestId() *string
	SetValue(v string) *PutDcdnKvResponseBody
	GetValue() *string
}

type PutDcdnKvResponseBody struct {
	// The length of the key.
	//
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 96ED3127-EC7A-57C5-AFA6-A689B24B2530
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of the key. If the value exceeds 256 characters in length, the first 100 characters and the last 100 characters are retained and other characters are not displayed.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *PutDcdnKvResponseBody) GetLength() *int32 {
	return s.Length
}

func (s *PutDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDcdnKvResponseBody) GetValue() *string {
	return s.Value
}

func (s *PutDcdnKvResponseBody) SetLength(v int32) *PutDcdnKvResponseBody {
	s.Length = &v
	return s
}

func (s *PutDcdnKvResponseBody) SetRequestId(v string) *PutDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDcdnKvResponseBody) SetValue(v string) *PutDcdnKvResponseBody {
	s.Value = &v
	return s
}

func (s *PutDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}
