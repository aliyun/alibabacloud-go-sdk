// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceIpWhitelistResponseBody) *CreateInstanceIpWhitelistResponse
	GetBody() *CreateInstanceIpWhitelistResponseBody
}

type CreateInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceIpWhitelistResponse) GetBody() *CreateInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *CreateInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *CreateInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceIpWhitelistResponse) SetStatusCode(v int32) *CreateInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponse) SetBody(v *CreateInstanceIpWhitelistResponseBody) *CreateInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
