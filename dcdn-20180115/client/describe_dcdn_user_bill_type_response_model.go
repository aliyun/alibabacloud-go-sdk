// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserBillTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserBillTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserBillTypeResponseBody) *DescribeDcdnUserBillTypeResponse
	GetBody() *DescribeDcdnUserBillTypeResponseBody
}

type DescribeDcdnUserBillTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserBillTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserBillTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserBillTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserBillTypeResponse) GetBody() *DescribeDcdnUserBillTypeResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserBillTypeResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserBillTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponse) SetStatusCode(v int32) *DescribeDcdnUserBillTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponse) SetBody(v *DescribeDcdnUserBillTypeResponseBody) *DescribeDcdnUserBillTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponse) Validate() error {
	return dara.Validate(s)
}
