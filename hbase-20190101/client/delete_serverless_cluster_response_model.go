// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServerlessClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServerlessClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServerlessClusterResponseBody) *DeleteServerlessClusterResponse
	GetBody() *DeleteServerlessClusterResponseBody
}

type DeleteServerlessClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerlessClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServerlessClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServerlessClusterResponse) GetBody() *DeleteServerlessClusterResponseBody {
	return s.Body
}

func (s *DeleteServerlessClusterResponse) SetHeaders(v map[string]*string) *DeleteServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerlessClusterResponse) SetStatusCode(v int32) *DeleteServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerlessClusterResponse) SetBody(v *DeleteServerlessClusterResponseBody) *DeleteServerlessClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteServerlessClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
