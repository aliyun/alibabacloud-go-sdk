// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateTagResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
}

type CreateTagResponseBody struct {
	// example:
	//
	// 7391272
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) SetId(v int64) *CreateTagResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
