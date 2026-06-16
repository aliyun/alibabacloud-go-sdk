// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchContextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchContextResponse
	GetStatusCode() *int32
	SetBody(v *SearchContextResponseBody) *SearchContextResponse
	GetBody() *SearchContextResponseBody
}

type SearchContextResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchContextResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchContextResponse) GoString() string {
	return s.String()
}

func (s *SearchContextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchContextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchContextResponse) GetBody() *SearchContextResponseBody {
	return s.Body
}

func (s *SearchContextResponse) SetHeaders(v map[string]*string) *SearchContextResponse {
	s.Headers = v
	return s
}

func (s *SearchContextResponse) SetStatusCode(v int32) *SearchContextResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchContextResponse) SetBody(v *SearchContextResponseBody) *SearchContextResponse {
	s.Body = v
	return s
}

func (s *SearchContextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
