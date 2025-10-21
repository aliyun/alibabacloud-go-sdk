// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSessionClusterResponseBody) *DeleteSessionClusterResponse
	GetBody() *DeleteSessionClusterResponseBody
}

type DeleteSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSessionClusterResponse) GetBody() *DeleteSessionClusterResponseBody {
	return s.Body
}

func (s *DeleteSessionClusterResponse) SetHeaders(v map[string]*string) *DeleteSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteSessionClusterResponse) SetStatusCode(v int32) *DeleteSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSessionClusterResponse) SetBody(v *DeleteSessionClusterResponseBody) *DeleteSessionClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteSessionClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
