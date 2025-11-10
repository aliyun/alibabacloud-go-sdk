// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceIpWhitelistResponseBody) *GetInstanceIpWhitelistResponse
	GetBody() *GetInstanceIpWhitelistResponseBody
}

type GetInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceIpWhitelistResponse) GetBody() *GetInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *GetInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *GetInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIpWhitelistResponse) SetStatusCode(v int32) *GetInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIpWhitelistResponse) SetBody(v *GetInstanceIpWhitelistResponseBody) *GetInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *GetInstanceIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
