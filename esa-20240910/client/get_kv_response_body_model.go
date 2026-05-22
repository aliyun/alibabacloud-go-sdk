// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetKvResponseBody
	GetRequestId() *string
	SetValue(v string) *GetKvResponseBody
	GetValue() *string
}

type GetKvResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKvResponseBody) GoString() string {
	return s.String()
}

func (s *GetKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKvResponseBody) GetValue() *string {
	return s.Value
}

func (s *GetKvResponseBody) SetRequestId(v string) *GetKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKvResponseBody) SetValue(v string) *GetKvResponseBody {
	s.Value = &v
	return s
}

func (s *GetKvResponseBody) Validate() error {
	return dara.Validate(s)
}
