// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDentryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDentryResponse
	GetStatusCode() *int32
	SetBody(v *QueryDentryResponseBody) *QueryDentryResponse
	GetBody() *QueryDentryResponseBody
}

type QueryDentryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDentryResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponse) GoString() string {
	return s.String()
}

func (s *QueryDentryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDentryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDentryResponse) GetBody() *QueryDentryResponseBody {
	return s.Body
}

func (s *QueryDentryResponse) SetHeaders(v map[string]*string) *QueryDentryResponse {
	s.Headers = v
	return s
}

func (s *QueryDentryResponse) SetStatusCode(v int32) *QueryDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDentryResponse) SetBody(v *QueryDentryResponseBody) *QueryDentryResponse {
	s.Body = v
	return s
}

func (s *QueryDentryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
