// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseOrderResponse
	GetStatusCode() *int32
	SetBody(v *CloseOrderResponseBody) *CloseOrderResponse
	GetBody() *CloseOrderResponseBody
}

type CloseOrderResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseOrderResponse) GetBody() *CloseOrderResponseBody {
	return s.Body
}

func (s *CloseOrderResponse) SetHeaders(v map[string]*string) *CloseOrderResponse {
	s.Headers = v
	return s
}

func (s *CloseOrderResponse) SetStatusCode(v int32) *CloseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseOrderResponse) SetBody(v *CloseOrderResponseBody) *CloseOrderResponse {
	s.Body = v
	return s
}

func (s *CloseOrderResponse) Validate() error {
	return dara.Validate(s)
}
