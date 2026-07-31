// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySelectOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySelectOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySelectOptionsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySelectOptionsResponseBody) *QuerySelectOptionsResponse
	GetBody() *QuerySelectOptionsResponseBody
}

type QuerySelectOptionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySelectOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySelectOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySelectOptionsResponse) GoString() string {
	return s.String()
}

func (s *QuerySelectOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySelectOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySelectOptionsResponse) GetBody() *QuerySelectOptionsResponseBody {
	return s.Body
}

func (s *QuerySelectOptionsResponse) SetHeaders(v map[string]*string) *QuerySelectOptionsResponse {
	s.Headers = v
	return s
}

func (s *QuerySelectOptionsResponse) SetStatusCode(v int32) *QuerySelectOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySelectOptionsResponse) SetBody(v *QuerySelectOptionsResponseBody) *QuerySelectOptionsResponse {
	s.Body = v
	return s
}

func (s *QuerySelectOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
