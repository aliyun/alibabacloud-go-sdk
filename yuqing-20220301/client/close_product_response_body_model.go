// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CloseProductResponseBody
	GetData() *int64
	SetRequestId(v string) *CloseProductResponseBody
	GetRequestId() *string
}

type CloseProductResponseBody struct {
	Data      *int64  `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloseProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseProductResponseBody) GoString() string {
	return s.String()
}

func (s *CloseProductResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CloseProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseProductResponseBody) SetData(v int64) *CloseProductResponseBody {
	s.Data = &v
	return s
}

func (s *CloseProductResponseBody) SetRequestId(v string) *CloseProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseProductResponseBody) Validate() error {
	return dara.Validate(s)
}
