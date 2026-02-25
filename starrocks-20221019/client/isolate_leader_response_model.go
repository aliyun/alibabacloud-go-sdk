// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsolateLeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsolateLeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsolateLeaderResponse
	GetStatusCode() *int32
	SetBody(v *IsolateLeaderResponseBody) *IsolateLeaderResponse
	GetBody() *IsolateLeaderResponseBody
}

type IsolateLeaderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsolateLeaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsolateLeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s IsolateLeaderResponse) GoString() string {
	return s.String()
}

func (s *IsolateLeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsolateLeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsolateLeaderResponse) GetBody() *IsolateLeaderResponseBody {
	return s.Body
}

func (s *IsolateLeaderResponse) SetHeaders(v map[string]*string) *IsolateLeaderResponse {
	s.Headers = v
	return s
}

func (s *IsolateLeaderResponse) SetStatusCode(v int32) *IsolateLeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *IsolateLeaderResponse) SetBody(v *IsolateLeaderResponseBody) *IsolateLeaderResponse {
	s.Body = v
	return s
}

func (s *IsolateLeaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
