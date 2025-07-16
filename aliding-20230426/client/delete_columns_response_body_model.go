// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteColumnsResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteColumnsResponseBody
	GetRequestId() *string
}

type DeleteColumnsResponseBody struct {
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

func (s DeleteColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteColumnsResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteColumnsResponseBody) SetId(v string) *DeleteColumnsResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteColumnsResponseBody) SetRequestId(v string) *DeleteColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteColumnsResponseBody) Validate() error {
	return dara.Validate(s)
}
