// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnKvNamespaceResponseBody) *DescribeDcdnKvNamespaceResponse
	GetBody() *DescribeDcdnKvNamespaceResponseBody
}

type DescribeDcdnKvNamespaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnKvNamespaceResponse) GetBody() *DescribeDcdnKvNamespaceResponseBody {
	return s.Body
}

func (s *DescribeDcdnKvNamespaceResponse) SetHeaders(v map[string]*string) *DescribeDcdnKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnKvNamespaceResponse) SetStatusCode(v int32) *DescribeDcdnKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponse) SetBody(v *DescribeDcdnKvNamespaceResponseBody) *DescribeDcdnKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnKvNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
