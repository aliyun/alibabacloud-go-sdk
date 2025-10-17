// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempAccessTokenIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TempAccessTokenIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TempAccessTokenIntlResponse
	GetStatusCode() *int32
	SetBody(v *TempAccessTokenIntlResponseBody) *TempAccessTokenIntlResponse
	GetBody() *TempAccessTokenIntlResponseBody
}

type TempAccessTokenIntlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TempAccessTokenIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TempAccessTokenIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s TempAccessTokenIntlResponse) GoString() string {
	return s.String()
}

func (s *TempAccessTokenIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TempAccessTokenIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TempAccessTokenIntlResponse) GetBody() *TempAccessTokenIntlResponseBody {
	return s.Body
}

func (s *TempAccessTokenIntlResponse) SetHeaders(v map[string]*string) *TempAccessTokenIntlResponse {
	s.Headers = v
	return s
}

func (s *TempAccessTokenIntlResponse) SetStatusCode(v int32) *TempAccessTokenIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *TempAccessTokenIntlResponse) SetBody(v *TempAccessTokenIntlResponseBody) *TempAccessTokenIntlResponse {
	s.Body = v
	return s
}

func (s *TempAccessTokenIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
