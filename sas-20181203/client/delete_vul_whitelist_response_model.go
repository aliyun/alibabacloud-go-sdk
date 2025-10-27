// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVulWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVulWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVulWhitelistResponseBody) *DeleteVulWhitelistResponse
	GetBody() *DeleteVulWhitelistResponseBody
}

type DeleteVulWhitelistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVulWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteVulWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVulWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVulWhitelistResponse) GetBody() *DeleteVulWhitelistResponseBody {
	return s.Body
}

func (s *DeleteVulWhitelistResponse) SetHeaders(v map[string]*string) *DeleteVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteVulWhitelistResponse) SetStatusCode(v int32) *DeleteVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVulWhitelistResponse) SetBody(v *DeleteVulWhitelistResponseBody) *DeleteVulWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteVulWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
