// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetAlipayUrlResponseBody
	GetData() *string
	SetRequestId(v string) *GetAlipayUrlResponseBody
	GetRequestId() *string
}

type GetAlipayUrlResponseBody struct {
	// example:
	//
	// https://
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlipayUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAlipayUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlipayUrlResponseBody) SetData(v string) *GetAlipayUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetAlipayUrlResponseBody) SetRequestId(v string) *GetAlipayUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlipayUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
