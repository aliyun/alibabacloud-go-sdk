// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListXTelephonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListXTelephonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListXTelephonesResponse
	GetStatusCode() *int32
	SetBody(v *ListXTelephonesResponseBody) *ListXTelephonesResponse
	GetBody() *ListXTelephonesResponseBody
}

type ListXTelephonesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListXTelephonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListXTelephonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponse) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListXTelephonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListXTelephonesResponse) GetBody() *ListXTelephonesResponseBody {
	return s.Body
}

func (s *ListXTelephonesResponse) SetHeaders(v map[string]*string) *ListXTelephonesResponse {
	s.Headers = v
	return s
}

func (s *ListXTelephonesResponse) SetStatusCode(v int32) *ListXTelephonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListXTelephonesResponse) SetBody(v *ListXTelephonesResponseBody) *ListXTelephonesResponse {
	s.Body = v
	return s
}

func (s *ListXTelephonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
