// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamPoolResponseBody) *DeleteIpamPoolResponse
	GetBody() *DeleteIpamPoolResponseBody
}

type DeleteIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamPoolResponse) GetBody() *DeleteIpamPoolResponseBody {
	return s.Body
}

func (s *DeleteIpamPoolResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolResponse) SetStatusCode(v int32) *DeleteIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolResponse) SetBody(v *DeleteIpamPoolResponseBody) *DeleteIpamPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
