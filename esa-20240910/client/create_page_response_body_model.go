// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreatePageResponseBody
	GetId() *int64
	SetRequestId(v string) *CreatePageResponseBody
	GetRequestId() *string
}

type CreatePageResponseBody struct {
	// The ID of the custom error page.[](~~2850223~~)
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePageResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreatePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePageResponseBody) SetId(v int64) *CreatePageResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePageResponseBody) SetRequestId(v string) *CreatePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePageResponseBody) Validate() error {
	return dara.Validate(s)
}
