// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteRowsResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteRowsResponseBody
	GetRequestId() *string
}

type DeleteRowsResponseBody struct {
	// example:
	//
	// stxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRowsResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteRowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRowsResponseBody) SetId(v string) *DeleteRowsResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteRowsResponseBody) SetRequestId(v string) *DeleteRowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRowsResponseBody) Validate() error {
	return dara.Validate(s)
}
