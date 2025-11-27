// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SimpleQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SimpleQueryResponse
	GetStatusCode() *int32
	SetBody(v *SimpleQueryResponseBody) *SimpleQueryResponse
	GetBody() *SimpleQueryResponseBody
}

type SimpleQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SimpleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SimpleQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryResponse) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SimpleQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SimpleQueryResponse) GetBody() *SimpleQueryResponseBody {
	return s.Body
}

func (s *SimpleQueryResponse) SetHeaders(v map[string]*string) *SimpleQueryResponse {
	s.Headers = v
	return s
}

func (s *SimpleQueryResponse) SetStatusCode(v int32) *SimpleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SimpleQueryResponse) SetBody(v *SimpleQueryResponseBody) *SimpleQueryResponse {
	s.Body = v
	return s
}

func (s *SimpleQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
