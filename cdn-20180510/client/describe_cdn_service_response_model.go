// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnServiceResponseBody) *DescribeCdnServiceResponse
	GetBody() *DescribeCdnServiceResponseBody
}

type DescribeCdnServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnServiceResponse) GetBody() *DescribeCdnServiceResponseBody {
	return s.Body
}

func (s *DescribeCdnServiceResponse) SetHeaders(v map[string]*string) *DescribeCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnServiceResponse) SetStatusCode(v int32) *DescribeCdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnServiceResponse) SetBody(v *DescribeCdnServiceResponseBody) *DescribeCdnServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnServiceResponse) Validate() error {
	return dara.Validate(s)
}
