// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkResponseBody) *DeleteNetworkResponse
	GetBody() *DeleteNetworkResponseBody
}

type DeleteNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkResponse) GetBody() *DeleteNetworkResponseBody {
	return s.Body
}

func (s *DeleteNetworkResponse) SetHeaders(v map[string]*string) *DeleteNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkResponse) SetStatusCode(v int32) *DeleteNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkResponse) SetBody(v *DeleteNetworkResponseBody) *DeleteNetworkResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
