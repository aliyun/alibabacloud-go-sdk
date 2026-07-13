// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceOssMountRamAuthorizeUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceOssMountRamAuthorizeUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceOssMountRamAuthorizeUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceOssMountRamAuthorizeUrlResponseBody) *GetInstanceOssMountRamAuthorizeUrlResponse
	GetBody() *GetInstanceOssMountRamAuthorizeUrlResponseBody
}

type GetInstanceOssMountRamAuthorizeUrlResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceOssMountRamAuthorizeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceOssMountRamAuthorizeUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceOssMountRamAuthorizeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) GetBody() *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	return s.Body
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) SetHeaders(v map[string]*string) *GetInstanceOssMountRamAuthorizeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) SetStatusCode(v int32) *GetInstanceOssMountRamAuthorizeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) SetBody(v *GetInstanceOssMountRamAuthorizeUrlResponseBody) *GetInstanceOssMountRamAuthorizeUrlResponse {
	s.Body = v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
