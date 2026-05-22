// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeletePageResponseBody
	GetId() *int64
	SetRequestId(v string) *DeletePageResponseBody
	GetRequestId() *string
}

type DeletePageResponseBody struct {
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

func (s DeletePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePageResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePageResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DeletePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePageResponseBody) SetId(v int64) *DeletePageResponseBody {
	s.Id = &v
	return s
}

func (s *DeletePageResponseBody) SetRequestId(v string) *DeletePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePageResponseBody) Validate() error {
	return dara.Validate(s)
}
