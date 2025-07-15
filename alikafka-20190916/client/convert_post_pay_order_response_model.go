// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPostPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertPostPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertPostPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *ConvertPostPayOrderResponseBody) *ConvertPostPayOrderResponse
	GetBody() *ConvertPostPayOrderResponseBody
}

type ConvertPostPayOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPostPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertPostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertPostPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertPostPayOrderResponse) GetBody() *ConvertPostPayOrderResponseBody {
	return s.Body
}

func (s *ConvertPostPayOrderResponse) SetHeaders(v map[string]*string) *ConvertPostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *ConvertPostPayOrderResponse) SetStatusCode(v int32) *ConvertPostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPostPayOrderResponse) SetBody(v *ConvertPostPayOrderResponseBody) *ConvertPostPayOrderResponse {
	s.Body = v
	return s
}

func (s *ConvertPostPayOrderResponse) Validate() error {
	return dara.Validate(s)
}
