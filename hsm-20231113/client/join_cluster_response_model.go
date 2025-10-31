// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinClusterResponse
	GetStatusCode() *int32
	SetBody(v *JoinClusterResponseBody) *JoinClusterResponse
	GetBody() *JoinClusterResponseBody
}

type JoinClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinClusterResponse) GoString() string {
	return s.String()
}

func (s *JoinClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinClusterResponse) GetBody() *JoinClusterResponseBody {
	return s.Body
}

func (s *JoinClusterResponse) SetHeaders(v map[string]*string) *JoinClusterResponse {
	s.Headers = v
	return s
}

func (s *JoinClusterResponse) SetStatusCode(v int32) *JoinClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinClusterResponse) SetBody(v *JoinClusterResponseBody) *JoinClusterResponse {
	s.Body = v
	return s
}

func (s *JoinClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
