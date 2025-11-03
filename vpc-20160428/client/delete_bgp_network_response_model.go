// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBgpNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBgpNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBgpNetworkResponseBody) *DeleteBgpNetworkResponse
	GetBody() *DeleteBgpNetworkResponseBody
}

type DeleteBgpNetworkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBgpNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBgpNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteBgpNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBgpNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBgpNetworkResponse) GetBody() *DeleteBgpNetworkResponseBody {
	return s.Body
}

func (s *DeleteBgpNetworkResponse) SetHeaders(v map[string]*string) *DeleteBgpNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteBgpNetworkResponse) SetStatusCode(v int32) *DeleteBgpNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBgpNetworkResponse) SetBody(v *DeleteBgpNetworkResponseBody) *DeleteBgpNetworkResponse {
	s.Body = v
	return s
}

func (s *DeleteBgpNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
