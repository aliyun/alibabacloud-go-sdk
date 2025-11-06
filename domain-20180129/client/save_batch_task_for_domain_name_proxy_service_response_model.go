// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForDomainNameProxyServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForDomainNameProxyServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForDomainNameProxyServiceResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForDomainNameProxyServiceResponseBody) *SaveBatchTaskForDomainNameProxyServiceResponse
	GetBody() *SaveBatchTaskForDomainNameProxyServiceResponseBody
}

type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForDomainNameProxyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForDomainNameProxyServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) GetBody() *SaveBatchTaskForDomainNameProxyServiceResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetStatusCode(v int32) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetBody(v *SaveBatchTaskForDomainNameProxyServiceResponseBody) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
