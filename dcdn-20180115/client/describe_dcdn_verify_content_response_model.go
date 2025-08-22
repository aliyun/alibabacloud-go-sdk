// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnVerifyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnVerifyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnVerifyContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnVerifyContentResponseBody) *DescribeDcdnVerifyContentResponse
	GetBody() *DescribeDcdnVerifyContentResponseBody
}

type DescribeDcdnVerifyContentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnVerifyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnVerifyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnVerifyContentResponse) GetBody() *DescribeDcdnVerifyContentResponseBody {
	return s.Body
}

func (s *DescribeDcdnVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeDcdnVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnVerifyContentResponse) SetStatusCode(v int32) *DescribeDcdnVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnVerifyContentResponse) SetBody(v *DescribeDcdnVerifyContentResponseBody) *DescribeDcdnVerifyContentResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnVerifyContentResponse) Validate() error {
	return dara.Validate(s)
}
