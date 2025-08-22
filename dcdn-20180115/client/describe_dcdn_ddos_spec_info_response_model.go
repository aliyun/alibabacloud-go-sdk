// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDdosSpecInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDdosSpecInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDdosSpecInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDdosSpecInfoResponseBody) *DescribeDcdnDdosSpecInfoResponse
	GetBody() *DescribeDcdnDdosSpecInfoResponseBody
}

type DescribeDcdnDdosSpecInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDdosSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDdosSpecInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosSpecInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDdosSpecInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDdosSpecInfoResponse) GetBody() *DescribeDcdnDdosSpecInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnDdosSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnDdosSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponse) SetStatusCode(v int32) *DescribeDcdnDdosSpecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponse) SetBody(v *DescribeDcdnDdosSpecInfoResponseBody) *DescribeDcdnDdosSpecInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDdosSpecInfoResponse) Validate() error {
	return dara.Validate(s)
}
