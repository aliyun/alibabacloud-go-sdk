// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListClientsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListClientsResponseBody) *ClinkListClientsResponse
	GetBody() *ClinkListClientsResponseBody
}

type ClinkListClientsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListClientsResponse) GoString() string {
	return s.String()
}

func (s *ClinkListClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListClientsResponse) GetBody() *ClinkListClientsResponseBody {
	return s.Body
}

func (s *ClinkListClientsResponse) SetHeaders(v map[string]*string) *ClinkListClientsResponse {
	s.Headers = v
	return s
}

func (s *ClinkListClientsResponse) SetStatusCode(v int32) *ClinkListClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListClientsResponse) SetBody(v *ClinkListClientsResponseBody) *ClinkListClientsResponse {
	s.Body = v
	return s
}

func (s *ClinkListClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
