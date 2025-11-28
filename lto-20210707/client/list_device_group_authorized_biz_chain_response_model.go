// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupAuthorizedBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceGroupAuthorizedBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceGroupAuthorizedBizChainResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceGroupAuthorizedBizChainResponseBody) *ListDeviceGroupAuthorizedBizChainResponse
	GetBody() *ListDeviceGroupAuthorizedBizChainResponseBody
}

type ListDeviceGroupAuthorizedBizChainResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceGroupAuthorizedBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceGroupAuthorizedBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupAuthorizedBizChainResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) GetBody() *ListDeviceGroupAuthorizedBizChainResponseBody {
	return s.Body
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) SetHeaders(v map[string]*string) *ListDeviceGroupAuthorizedBizChainResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) SetStatusCode(v int32) *ListDeviceGroupAuthorizedBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) SetBody(v *ListDeviceGroupAuthorizedBizChainResponseBody) *ListDeviceGroupAuthorizedBizChainResponse {
	s.Body = v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
