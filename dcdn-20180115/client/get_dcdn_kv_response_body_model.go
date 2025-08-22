// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDcdnKvResponseBody
	GetRequestId() *string
	SetValue(v string) *GetDcdnKvResponseBody
	GetValue() *string
}

type GetDcdnKvResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *GetDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDcdnKvResponseBody) GetValue() *string {
	return s.Value
}

func (s *GetDcdnKvResponseBody) SetRequestId(v string) *GetDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDcdnKvResponseBody) SetValue(v string) *GetDcdnKvResponseBody {
	s.Value = &v
	return s
}

func (s *GetDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}
