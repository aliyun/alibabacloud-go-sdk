// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateResourceResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateResourceResponseBody
	GetRequestId() *string
}

type CreateResourceResponseBody struct {
	// The ID of the file resource.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// A5B97987-66EA-5563-9599-A2752292XXXX
	//
	// example:
	//
	// The ID of the file resource.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceResponseBody) SetId(v int64) *CreateResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
