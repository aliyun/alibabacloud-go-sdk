// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventIspResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosEventIspResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosEventIspResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosEventIspResponseBody) *DescribeDDosEventIspResponse
	GetBody() *DescribeDDosEventIspResponseBody
}

type DescribeDDosEventIspResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosEventIspResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosEventIspResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosEventIspResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosEventIspResponse) GetBody() *DescribeDDosEventIspResponseBody {
	return s.Body
}

func (s *DescribeDDosEventIspResponse) SetHeaders(v map[string]*string) *DescribeDDosEventIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventIspResponse) SetStatusCode(v int32) *DescribeDDosEventIspResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosEventIspResponse) SetBody(v *DescribeDDosEventIspResponseBody) *DescribeDDosEventIspResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosEventIspResponse) Validate() error {
	return dara.Validate(s)
}
