// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGoodsShippingNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGoodsShippingNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGoodsShippingNoticeResponse
	GetStatusCode() *int32
	SetBody(v *GoodsShippingNoticeCreateResult) *CreateGoodsShippingNoticeResponse
	GetBody() *GoodsShippingNoticeCreateResult
}

type CreateGoodsShippingNoticeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GoodsShippingNoticeCreateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGoodsShippingNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGoodsShippingNoticeResponse) GoString() string {
	return s.String()
}

func (s *CreateGoodsShippingNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGoodsShippingNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGoodsShippingNoticeResponse) GetBody() *GoodsShippingNoticeCreateResult {
	return s.Body
}

func (s *CreateGoodsShippingNoticeResponse) SetHeaders(v map[string]*string) *CreateGoodsShippingNoticeResponse {
	s.Headers = v
	return s
}

func (s *CreateGoodsShippingNoticeResponse) SetStatusCode(v int32) *CreateGoodsShippingNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGoodsShippingNoticeResponse) SetBody(v *GoodsShippingNoticeCreateResult) *CreateGoodsShippingNoticeResponse {
	s.Body = v
	return s
}

func (s *CreateGoodsShippingNoticeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
