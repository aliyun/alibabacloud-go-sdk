// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainExtensionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainExtensionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainExtensionAttributeResponseBody) *DescribeDomainExtensionAttributeResponse
	GetBody() *DescribeDomainExtensionAttributeResponseBody
}

type DescribeDomainExtensionAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainExtensionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainExtensionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainExtensionAttributeResponse) GetBody() *DescribeDomainExtensionAttributeResponseBody {
	return s.Body
}

func (s *DescribeDomainExtensionAttributeResponse) SetHeaders(v map[string]*string) *DescribeDomainExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainExtensionAttributeResponse) SetStatusCode(v int32) *DescribeDomainExtensionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponse) SetBody(v *DescribeDomainExtensionAttributeResponseBody) *DescribeDomainExtensionAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainExtensionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
