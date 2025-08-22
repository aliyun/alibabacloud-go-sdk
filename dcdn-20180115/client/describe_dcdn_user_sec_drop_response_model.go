// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserSecDropResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserSecDropResponseBody) *DescribeDcdnUserSecDropResponse
	GetBody() *DescribeDcdnUserSecDropResponseBody
}

type DescribeDcdnUserSecDropResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserSecDropResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserSecDropResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserSecDropResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserSecDropResponse) GetBody() *DescribeDcdnUserSecDropResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserSecDropResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserSecDropResponse) SetStatusCode(v int32) *DescribeDcdnUserSecDropResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponse) SetBody(v *DescribeDcdnUserSecDropResponseBody) *DescribeDcdnUserSecDropResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserSecDropResponse) Validate() error {
	return dara.Validate(s)
}
