// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteProductResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteProductResponseBody
	GetRequestId() *string
}

type DeleteProductResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2dc580ca52ce43f59f890eaec37c8dad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProductResponseBody) SetData(v bool) *DeleteProductResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductResponseBody) Validate() error {
	return dara.Validate(s)
}
