// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAlipayUrlResponseBodyData) *GetAlipayUrlResponseBody
	GetData() *GetAlipayUrlResponseBodyData
	SetRequestId(v string) *GetAlipayUrlResponseBody
	GetRequestId() *string
}

type GetAlipayUrlResponseBody struct {
	Data *GetAlipayUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-11402910xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAlipayUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponseBody) GetData() *GetAlipayUrlResponseBodyData {
	return s.Data
}

func (s *GetAlipayUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlipayUrlResponseBody) SetData(v *GetAlipayUrlResponseBodyData) *GetAlipayUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAlipayUrlResponseBody) SetRequestId(v string) *GetAlipayUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlipayUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlipayUrlResponseBodyData struct {
	// example:
	//
	// xxsdfasfw
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
	QrUrl *string `json:"qrUrl,omitempty" xml:"qrUrl,omitempty"`
}

func (s GetAlipayUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetAlipayUrlResponseBodyData) GetQrUrl() *string {
	return s.QrUrl
}

func (s *GetAlipayUrlResponseBodyData) SetCode(v string) *GetAlipayUrlResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAlipayUrlResponseBodyData) SetQrUrl(v string) *GetAlipayUrlResponseBodyData {
	s.QrUrl = &v
	return s
}

func (s *GetAlipayUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
