// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserDomainsByFuncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserDomainsByFuncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserDomainsByFuncResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserDomainsByFuncResponseBody) *DescribeCdnUserDomainsByFuncResponse
	GetBody() *DescribeCdnUserDomainsByFuncResponseBody
}

type DescribeCdnUserDomainsByFuncResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserDomainsByFuncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserDomainsByFuncResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserDomainsByFuncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserDomainsByFuncResponse) GetBody() *DescribeCdnUserDomainsByFuncResponseBody {
	return s.Body
}

func (s *DescribeCdnUserDomainsByFuncResponse) SetHeaders(v map[string]*string) *DescribeCdnUserDomainsByFuncResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponse) SetStatusCode(v int32) *DescribeCdnUserDomainsByFuncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponse) SetBody(v *DescribeCdnUserDomainsByFuncResponseBody) *DescribeCdnUserDomainsByFuncResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponse) Validate() error {
	return dara.Validate(s)
}
