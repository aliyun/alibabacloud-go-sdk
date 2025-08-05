// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteQuotaPlanResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteQuotaPlanResponseBody
	GetRequestId() *string
}

type DeleteQuotaPlanResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQuotaPlanResponseBody) SetData(v string) *DeleteQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQuotaPlanResponseBody) SetRequestId(v string) *DeleteQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
