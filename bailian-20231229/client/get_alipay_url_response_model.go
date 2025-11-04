// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlipayUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlipayUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAlipayUrlResponseBody) *GetAlipayUrlResponse
	GetBody() *GetAlipayUrlResponseBody
}

type GetAlipayUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlipayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlipayUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlipayUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlipayUrlResponse) GetBody() *GetAlipayUrlResponseBody {
	return s.Body
}

func (s *GetAlipayUrlResponse) SetHeaders(v map[string]*string) *GetAlipayUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAlipayUrlResponse) SetStatusCode(v int32) *GetAlipayUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlipayUrlResponse) SetBody(v *GetAlipayUrlResponseBody) *GetAlipayUrlResponse {
	s.Body = v
	return s
}

func (s *GetAlipayUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
