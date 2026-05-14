// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDeleteClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDeleteClientResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDeleteClientResponseBody) *ClinkDeleteClientResponse
	GetBody() *ClinkDeleteClientResponseBody
}

type ClinkDeleteClientResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDeleteClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDeleteClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteClientResponse) GoString() string {
	return s.String()
}

func (s *ClinkDeleteClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDeleteClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDeleteClientResponse) GetBody() *ClinkDeleteClientResponseBody {
	return s.Body
}

func (s *ClinkDeleteClientResponse) SetHeaders(v map[string]*string) *ClinkDeleteClientResponse {
	s.Headers = v
	return s
}

func (s *ClinkDeleteClientResponse) SetStatusCode(v int32) *ClinkDeleteClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDeleteClientResponse) SetBody(v *ClinkDeleteClientResponseBody) *ClinkDeleteClientResponse {
	s.Body = v
	return s
}

func (s *ClinkDeleteClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
