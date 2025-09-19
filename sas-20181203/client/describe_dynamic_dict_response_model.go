// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDynamicDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDynamicDictResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDynamicDictResponseBody) *DescribeDynamicDictResponse
	GetBody() *DescribeDynamicDictResponseBody
}

type DescribeDynamicDictResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDynamicDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDynamicDictResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictResponse) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDynamicDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDynamicDictResponse) GetBody() *DescribeDynamicDictResponseBody {
	return s.Body
}

func (s *DescribeDynamicDictResponse) SetHeaders(v map[string]*string) *DescribeDynamicDictResponse {
	s.Headers = v
	return s
}

func (s *DescribeDynamicDictResponse) SetStatusCode(v int32) *DescribeDynamicDictResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDynamicDictResponse) SetBody(v *DescribeDynamicDictResponseBody) *DescribeDynamicDictResponse {
	s.Body = v
	return s
}

func (s *DescribeDynamicDictResponse) Validate() error {
	return dara.Validate(s)
}
