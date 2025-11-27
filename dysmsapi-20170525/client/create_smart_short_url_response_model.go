// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartShortUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmartShortUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmartShortUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmartShortUrlResponseBody) *CreateSmartShortUrlResponse
	GetBody() *CreateSmartShortUrlResponseBody
}

type CreateSmartShortUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartShortUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartShortUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmartShortUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmartShortUrlResponse) GetBody() *CreateSmartShortUrlResponseBody {
	return s.Body
}

func (s *CreateSmartShortUrlResponse) SetHeaders(v map[string]*string) *CreateSmartShortUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartShortUrlResponse) SetStatusCode(v int32) *CreateSmartShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartShortUrlResponse) SetBody(v *CreateSmartShortUrlResponseBody) *CreateSmartShortUrlResponse {
	s.Body = v
	return s
}

func (s *CreateSmartShortUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
