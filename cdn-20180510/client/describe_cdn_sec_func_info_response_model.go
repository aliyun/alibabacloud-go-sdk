// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSecFuncInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnSecFuncInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnSecFuncInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnSecFuncInfoResponseBody) *DescribeCdnSecFuncInfoResponse
	GetBody() *DescribeCdnSecFuncInfoResponseBody
}

type DescribeCdnSecFuncInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnSecFuncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnSecFuncInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSecFuncInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSecFuncInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnSecFuncInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnSecFuncInfoResponse) GetBody() *DescribeCdnSecFuncInfoResponseBody {
	return s.Body
}

func (s *DescribeCdnSecFuncInfoResponse) SetHeaders(v map[string]*string) *DescribeCdnSecFuncInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSecFuncInfoResponse) SetStatusCode(v int32) *DescribeCdnSecFuncInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponse) SetBody(v *DescribeCdnSecFuncInfoResponseBody) *DescribeCdnSecFuncInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnSecFuncInfoResponse) Validate() error {
	return dara.Validate(s)
}
