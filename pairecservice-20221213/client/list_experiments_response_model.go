// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperimentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperimentsResponse
	GetStatusCode() *int32
	SetBody(v *ListExperimentsResponseBody) *ListExperimentsResponse
	GetBody() *ListExperimentsResponseBody
}

type ListExperimentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperimentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperimentsResponse) GetBody() *ListExperimentsResponseBody {
	return s.Body
}

func (s *ListExperimentsResponse) SetHeaders(v map[string]*string) *ListExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentsResponse) SetStatusCode(v int32) *ListExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentsResponse) SetBody(v *ListExperimentsResponseBody) *ListExperimentsResponse {
	s.Body = v
	return s
}

func (s *ListExperimentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
