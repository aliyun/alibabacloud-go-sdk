// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TermQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TermQueryResponse
	GetStatusCode() *int32
	SetBody(v *TermQueryResponseBody) *TermQueryResponse
	GetBody() *TermQueryResponseBody
}

type TermQueryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TermQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TermQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TermQueryResponse) GoString() string {
	return s.String()
}

func (s *TermQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TermQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TermQueryResponse) GetBody() *TermQueryResponseBody {
	return s.Body
}

func (s *TermQueryResponse) SetHeaders(v map[string]*string) *TermQueryResponse {
	s.Headers = v
	return s
}

func (s *TermQueryResponse) SetStatusCode(v int32) *TermQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TermQueryResponse) SetBody(v *TermQueryResponseBody) *TermQueryResponse {
	s.Body = v
	return s
}

func (s *TermQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
