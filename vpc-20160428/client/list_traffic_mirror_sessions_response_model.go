// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficMirrorSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficMirrorSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficMirrorSessionsResponseBody) *ListTrafficMirrorSessionsResponse
	GetBody() *ListTrafficMirrorSessionsResponseBody
}

type ListTrafficMirrorSessionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficMirrorSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficMirrorSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficMirrorSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficMirrorSessionsResponse) GetBody() *ListTrafficMirrorSessionsResponseBody {
	return s.Body
}

func (s *ListTrafficMirrorSessionsResponse) SetHeaders(v map[string]*string) *ListTrafficMirrorSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficMirrorSessionsResponse) SetStatusCode(v int32) *ListTrafficMirrorSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficMirrorSessionsResponse) SetBody(v *ListTrafficMirrorSessionsResponseBody) *ListTrafficMirrorSessionsResponse {
	s.Body = v
	return s
}

func (s *ListTrafficMirrorSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
