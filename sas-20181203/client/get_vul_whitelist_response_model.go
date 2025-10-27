// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *GetVulWhitelistResponseBody) *GetVulWhitelistResponse
	GetBody() *GetVulWhitelistResponseBody
}

type GetVulWhitelistResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *GetVulWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulWhitelistResponse) GetBody() *GetVulWhitelistResponseBody {
	return s.Body
}

func (s *GetVulWhitelistResponse) SetHeaders(v map[string]*string) *GetVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *GetVulWhitelistResponse) SetStatusCode(v int32) *GetVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulWhitelistResponse) SetBody(v *GetVulWhitelistResponseBody) *GetVulWhitelistResponse {
	s.Body = v
	return s
}

func (s *GetVulWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
