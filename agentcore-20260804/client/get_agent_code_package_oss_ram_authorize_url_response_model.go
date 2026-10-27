// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageOssRamAuthorizeUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentCodePackageOssRamAuthorizeUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentCodePackageOssRamAuthorizeUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) *GetAgentCodePackageOssRamAuthorizeUrlResponse
	GetBody() *GetAgentCodePackageOssRamAuthorizeUrlResponseBody
}

type GetAgentCodePackageOssRamAuthorizeUrlResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentCodePackageOssRamAuthorizeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageOssRamAuthorizeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) GetBody() *GetAgentCodePackageOssRamAuthorizeUrlResponseBody {
	return s.Body
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) SetHeaders(v map[string]*string) *GetAgentCodePackageOssRamAuthorizeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) SetStatusCode(v int32) *GetAgentCodePackageOssRamAuthorizeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) SetBody(v *GetAgentCodePackageOssRamAuthorizeUrlResponseBody) *GetAgentCodePackageOssRamAuthorizeUrlResponse {
	s.Body = v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
