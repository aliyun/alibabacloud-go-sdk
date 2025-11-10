// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceIpWhitelistResponseBody) *DeleteInstanceIpWhitelistResponse
	GetBody() *DeleteInstanceIpWhitelistResponseBody
}

type DeleteInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceIpWhitelistResponse) GetBody() *DeleteInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *DeleteInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *DeleteInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceIpWhitelistResponse) SetStatusCode(v int32) *DeleteInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponse) SetBody(v *DeleteInstanceIpWhitelistResponseBody) *DeleteInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
