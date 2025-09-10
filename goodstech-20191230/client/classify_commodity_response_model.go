// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClassifyCommodityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClassifyCommodityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClassifyCommodityResponse
	GetStatusCode() *int32
	SetBody(v *ClassifyCommodityResponseBody) *ClassifyCommodityResponse
	GetBody() *ClassifyCommodityResponseBody
}

type ClassifyCommodityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClassifyCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClassifyCommodityResponse) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityResponse) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClassifyCommodityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClassifyCommodityResponse) GetBody() *ClassifyCommodityResponseBody {
	return s.Body
}

func (s *ClassifyCommodityResponse) SetHeaders(v map[string]*string) *ClassifyCommodityResponse {
	s.Headers = v
	return s
}

func (s *ClassifyCommodityResponse) SetStatusCode(v int32) *ClassifyCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *ClassifyCommodityResponse) SetBody(v *ClassifyCommodityResponseBody) *ClassifyCommodityResponse {
	s.Body = v
	return s
}

func (s *ClassifyCommodityResponse) Validate() error {
	return dara.Validate(s)
}
