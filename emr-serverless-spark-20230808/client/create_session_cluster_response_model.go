// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateSessionClusterResponseBody) *CreateSessionClusterResponse
	GetBody() *CreateSessionClusterResponseBody
}

type CreateSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSessionClusterResponse) GetBody() *CreateSessionClusterResponseBody {
	return s.Body
}

func (s *CreateSessionClusterResponse) SetHeaders(v map[string]*string) *CreateSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionClusterResponse) SetStatusCode(v int32) *CreateSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSessionClusterResponse) SetBody(v *CreateSessionClusterResponseBody) *CreateSessionClusterResponse {
	s.Body = v
	return s
}

func (s *CreateSessionClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
