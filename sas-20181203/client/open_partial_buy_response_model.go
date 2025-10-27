// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPartialBuyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenPartialBuyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenPartialBuyResponse
	GetStatusCode() *int32
	SetBody(v *OpenPartialBuyResponseBody) *OpenPartialBuyResponse
	GetBody() *OpenPartialBuyResponseBody
}

type OpenPartialBuyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenPartialBuyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenPartialBuyResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenPartialBuyResponse) GoString() string {
	return s.String()
}

func (s *OpenPartialBuyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenPartialBuyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenPartialBuyResponse) GetBody() *OpenPartialBuyResponseBody {
	return s.Body
}

func (s *OpenPartialBuyResponse) SetHeaders(v map[string]*string) *OpenPartialBuyResponse {
	s.Headers = v
	return s
}

func (s *OpenPartialBuyResponse) SetStatusCode(v int32) *OpenPartialBuyResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenPartialBuyResponse) SetBody(v *OpenPartialBuyResponseBody) *OpenPartialBuyResponse {
	s.Body = v
	return s
}

func (s *OpenPartialBuyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
