// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamResponseBody) *DeleteIpamResponse
	GetBody() *DeleteIpamResponseBody
}

type DeleteIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamResponse) GetBody() *DeleteIpamResponseBody {
	return s.Body
}

func (s *DeleteIpamResponse) SetHeaders(v map[string]*string) *DeleteIpamResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamResponse) SetStatusCode(v int32) *DeleteIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamResponse) SetBody(v *DeleteIpamResponseBody) *DeleteIpamResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
