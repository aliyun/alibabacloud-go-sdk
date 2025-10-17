// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse
	GetBody() *DeleteDBClusterResponseBody
}

type DeleteDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBClusterResponse) GetBody() *DeleteDBClusterResponseBody {
	return s.Body
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetStatusCode(v int32) *DeleteDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteDBClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
