// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentExperimentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentExperimentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentExperimentsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentExperimentsResponseBody) *ListRecentExperimentsResponse
	GetBody() *ListRecentExperimentsResponseBody
}

type ListRecentExperimentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentExperimentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentExperimentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentExperimentsResponse) GetBody() *ListRecentExperimentsResponseBody {
	return s.Body
}

func (s *ListRecentExperimentsResponse) SetHeaders(v map[string]*string) *ListRecentExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentExperimentsResponse) SetStatusCode(v int32) *ListRecentExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentExperimentsResponse) SetBody(v *ListRecentExperimentsResponseBody) *ListRecentExperimentsResponse {
	s.Body = v
	return s
}

func (s *ListRecentExperimentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
