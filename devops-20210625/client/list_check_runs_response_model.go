// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckRunsResponseBody) *ListCheckRunsResponse
	GetBody() *ListCheckRunsResponseBody
}

type ListCheckRunsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponse) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckRunsResponse) GetBody() *ListCheckRunsResponseBody {
	return s.Body
}

func (s *ListCheckRunsResponse) SetHeaders(v map[string]*string) *ListCheckRunsResponse {
	s.Headers = v
	return s
}

func (s *ListCheckRunsResponse) SetStatusCode(v int32) *ListCheckRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckRunsResponse) SetBody(v *ListCheckRunsResponseBody) *ListCheckRunsResponse {
	s.Body = v
	return s
}

func (s *ListCheckRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
