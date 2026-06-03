// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDomainNameProxyServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForDomainNameProxyServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForDomainNameProxyServiceResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForDomainNameProxyServiceResponseBody) *SaveSingleTaskForDomainNameProxyServiceResponse
	GetBody() *SaveSingleTaskForDomainNameProxyServiceResponseBody
}

type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForDomainNameProxyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForDomainNameProxyServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) GetBody() *SaveSingleTaskForDomainNameProxyServiceResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetStatusCode(v int32) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetBody(v *SaveSingleTaskForDomainNameProxyServiceResponseBody) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
