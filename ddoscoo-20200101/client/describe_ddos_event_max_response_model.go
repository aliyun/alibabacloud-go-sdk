// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventMaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosEventMaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosEventMaxResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosEventMaxResponseBody) *DescribeDDosEventMaxResponse
	GetBody() *DescribeDDosEventMaxResponseBody
}

type DescribeDDosEventMaxResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosEventMaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosEventMaxResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventMaxResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosEventMaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosEventMaxResponse) GetBody() *DescribeDDosEventMaxResponseBody {
	return s.Body
}

func (s *DescribeDDosEventMaxResponse) SetHeaders(v map[string]*string) *DescribeDDosEventMaxResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventMaxResponse) SetStatusCode(v int32) *DescribeDDosEventMaxResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosEventMaxResponse) SetBody(v *DescribeDDosEventMaxResponseBody) *DescribeDDosEventMaxResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosEventMaxResponse) Validate() error {
	return dara.Validate(s)
}
