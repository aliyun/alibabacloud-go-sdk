// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsAuthorizationLetterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsAuthorizationLetterResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsAuthorizationLetterResponseBody) *QuerySmsAuthorizationLetterResponse
	GetBody() *QuerySmsAuthorizationLetterResponseBody
}

type QuerySmsAuthorizationLetterResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsAuthorizationLetterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsAuthorizationLetterResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsAuthorizationLetterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsAuthorizationLetterResponse) GetBody() *QuerySmsAuthorizationLetterResponseBody {
	return s.Body
}

func (s *QuerySmsAuthorizationLetterResponse) SetHeaders(v map[string]*string) *QuerySmsAuthorizationLetterResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsAuthorizationLetterResponse) SetStatusCode(v int32) *QuerySmsAuthorizationLetterResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponse) SetBody(v *QuerySmsAuthorizationLetterResponseBody) *QuerySmsAuthorizationLetterResponse {
	s.Body = v
	return s
}

func (s *QuerySmsAuthorizationLetterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
