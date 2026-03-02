// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeletePdpConfigResponseBody) *DeletePdpConfigResponse
	GetBody() *DeletePdpConfigResponseBody
}

type DeletePdpConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePdpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpConfigResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpConfigResponse) GetBody() *DeletePdpConfigResponseBody {
	return s.Body
}

func (s *DeletePdpConfigResponse) SetHeaders(v map[string]*string) *DeletePdpConfigResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpConfigResponse) SetStatusCode(v int32) *DeletePdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpConfigResponse) SetBody(v *DeletePdpConfigResponseBody) *DeletePdpConfigResponse {
	s.Body = v
	return s
}

func (s *DeletePdpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
