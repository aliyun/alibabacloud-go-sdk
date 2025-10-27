// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPhoneResponse
	GetStatusCode() *int32
	SetBody(v *BindPhoneResponseBody) *BindPhoneResponse
	GetBody() *BindPhoneResponseBody
}

type BindPhoneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPhoneResponse) GoString() string {
	return s.String()
}

func (s *BindPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPhoneResponse) GetBody() *BindPhoneResponseBody {
	return s.Body
}

func (s *BindPhoneResponse) SetHeaders(v map[string]*string) *BindPhoneResponse {
	s.Headers = v
	return s
}

func (s *BindPhoneResponse) SetStatusCode(v int32) *BindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPhoneResponse) SetBody(v *BindPhoneResponseBody) *BindPhoneResponse {
	s.Body = v
	return s
}

func (s *BindPhoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
