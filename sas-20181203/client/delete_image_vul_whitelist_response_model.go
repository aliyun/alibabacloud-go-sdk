// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageVulWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageVulWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageVulWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageVulWhitelistResponseBody) *DeleteImageVulWhitelistResponse
	GetBody() *DeleteImageVulWhitelistResponseBody
}

type DeleteImageVulWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageVulWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageVulWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageVulWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageVulWhitelistResponse) GetBody() *DeleteImageVulWhitelistResponseBody {
	return s.Body
}

func (s *DeleteImageVulWhitelistResponse) SetHeaders(v map[string]*string) *DeleteImageVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageVulWhitelistResponse) SetStatusCode(v int32) *DeleteImageVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageVulWhitelistResponse) SetBody(v *DeleteImageVulWhitelistResponseBody) *DeleteImageVulWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteImageVulWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
