// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeDeviceGroupBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeDeviceGroupBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeDeviceGroupBizChainResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeDeviceGroupBizChainResponseBody) *AuthorizeDeviceGroupBizChainResponse
	GetBody() *AuthorizeDeviceGroupBizChainResponseBody
}

type AuthorizeDeviceGroupBizChainResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeDeviceGroupBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeDeviceGroupBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeDeviceGroupBizChainResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeDeviceGroupBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeDeviceGroupBizChainResponse) GetBody() *AuthorizeDeviceGroupBizChainResponseBody {
	return s.Body
}

func (s *AuthorizeDeviceGroupBizChainResponse) SetHeaders(v map[string]*string) *AuthorizeDeviceGroupBizChainResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponse) SetStatusCode(v int32) *AuthorizeDeviceGroupBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponse) SetBody(v *AuthorizeDeviceGroupBizChainResponseBody) *AuthorizeDeviceGroupBizChainResponse {
	s.Body = v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
