// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateListResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateListResponseBody
	GetRequestId() *string
}

type CreateListResponseBody struct {
	// The ID of the custom list.[](~~2850217~~)
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateListResponseBody) SetId(v int64) *CreateListResponseBody {
	s.Id = &v
	return s
}

func (s *CreateListResponseBody) SetRequestId(v string) *CreateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateListResponseBody) Validate() error {
	return dara.Validate(s)
}
