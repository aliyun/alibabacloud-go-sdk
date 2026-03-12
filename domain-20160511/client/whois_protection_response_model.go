// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhoisProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WhoisProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WhoisProtectionResponse
	GetStatusCode() *int32
	SetBody(v *WhoisProtectionResponseBody) *WhoisProtectionResponse
	GetBody() *WhoisProtectionResponseBody
}

type WhoisProtectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WhoisProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WhoisProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s WhoisProtectionResponse) GoString() string {
	return s.String()
}

func (s *WhoisProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WhoisProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WhoisProtectionResponse) GetBody() *WhoisProtectionResponseBody {
	return s.Body
}

func (s *WhoisProtectionResponse) SetHeaders(v map[string]*string) *WhoisProtectionResponse {
	s.Headers = v
	return s
}

func (s *WhoisProtectionResponse) SetStatusCode(v int32) *WhoisProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *WhoisProtectionResponse) SetBody(v *WhoisProtectionResponseBody) *WhoisProtectionResponse {
	s.Body = v
	return s
}

func (s *WhoisProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
