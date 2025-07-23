// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOrderRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOrderRequestResponse
	GetStatusCode() *int32
	SetBody(v *CancelOrderRequestResponseBody) *CancelOrderRequestResponse
	GetBody() *CancelOrderRequestResponseBody
}

type CancelOrderRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRequestResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOrderRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOrderRequestResponse) GetBody() *CancelOrderRequestResponseBody {
	return s.Body
}

func (s *CancelOrderRequestResponse) SetHeaders(v map[string]*string) *CancelOrderRequestResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderRequestResponse) SetStatusCode(v int32) *CancelOrderRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderRequestResponse) SetBody(v *CancelOrderRequestResponseBody) *CancelOrderRequestResponse {
	s.Body = v
	return s
}

func (s *CancelOrderRequestResponse) Validate() error {
	return dara.Validate(s)
}
