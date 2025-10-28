// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourceDriftsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackResourceDriftsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackResourceDriftsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackResourceDriftsResponseBody) *ListStackResourceDriftsResponse
	GetBody() *ListStackResourceDriftsResponseBody
}

type ListStackResourceDriftsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackResourceDriftsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackResourceDriftsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackResourceDriftsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackResourceDriftsResponse) GetBody() *ListStackResourceDriftsResponseBody {
	return s.Body
}

func (s *ListStackResourceDriftsResponse) SetHeaders(v map[string]*string) *ListStackResourceDriftsResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourceDriftsResponse) SetStatusCode(v int32) *ListStackResourceDriftsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackResourceDriftsResponse) SetBody(v *ListStackResourceDriftsResponseBody) *ListStackResourceDriftsResponse {
	s.Body = v
	return s
}

func (s *ListStackResourceDriftsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
