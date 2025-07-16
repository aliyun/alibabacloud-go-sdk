// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnTypesResponseBody) *DescribeCdnTypesResponse
	GetBody() *DescribeCdnTypesResponseBody
}

type DescribeCdnTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnTypesResponse) GetBody() *DescribeCdnTypesResponseBody {
	return s.Body
}

func (s *DescribeCdnTypesResponse) SetHeaders(v map[string]*string) *DescribeCdnTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnTypesResponse) SetStatusCode(v int32) *DescribeCdnTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnTypesResponse) SetBody(v *DescribeCdnTypesResponseBody) *DescribeCdnTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnTypesResponse) Validate() error {
	return dara.Validate(s)
}
