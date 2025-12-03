// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *OpenProductResponseBody
	GetId() *int64
	SetRequestId(v string) *OpenProductResponseBody
	GetRequestId() *string
}

type OpenProductResponseBody struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenProductResponseBody) GoString() string {
	return s.String()
}

func (s *OpenProductResponseBody) GetId() *int64 {
	return s.Id
}

func (s *OpenProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenProductResponseBody) SetId(v int64) *OpenProductResponseBody {
	s.Id = &v
	return s
}

func (s *OpenProductResponseBody) SetRequestId(v string) *OpenProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenProductResponseBody) Validate() error {
	return dara.Validate(s)
}
