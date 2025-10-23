// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpfilterByEdmIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpfilterByEdmIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpfilterByEdmIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpfilterByEdmIdResponseBody) *DeleteIpfilterByEdmIdResponse
	GetBody() *DeleteIpfilterByEdmIdResponseBody
}

type DeleteIpfilterByEdmIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpfilterByEdmIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpfilterByEdmIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpfilterByEdmIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpfilterByEdmIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpfilterByEdmIdResponse) GetBody() *DeleteIpfilterByEdmIdResponseBody {
	return s.Body
}

func (s *DeleteIpfilterByEdmIdResponse) SetHeaders(v map[string]*string) *DeleteIpfilterByEdmIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpfilterByEdmIdResponse) SetStatusCode(v int32) *DeleteIpfilterByEdmIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpfilterByEdmIdResponse) SetBody(v *DeleteIpfilterByEdmIdResponseBody) *DeleteIpfilterByEdmIdResponse {
	s.Body = v
	return s
}

func (s *DeleteIpfilterByEdmIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
