// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDynamicDictResponseBody
	GetRequestId() *string
}

type CreateDynamicDictResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EACE89CB-F32B-5A85-9242-D474A2ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDynamicDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicDictResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDynamicDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDynamicDictResponseBody) SetRequestId(v string) *CreateDynamicDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDynamicDictResponseBody) Validate() error {
	return dara.Validate(s)
}
