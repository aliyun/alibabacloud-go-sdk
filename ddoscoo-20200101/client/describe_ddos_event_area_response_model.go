// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAreaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosEventAreaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosEventAreaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosEventAreaResponseBody) *DescribeDDosEventAreaResponse
	GetBody() *DescribeDDosEventAreaResponseBody
}

type DescribeDDosEventAreaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosEventAreaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosEventAreaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAreaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosEventAreaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosEventAreaResponse) GetBody() *DescribeDDosEventAreaResponseBody {
	return s.Body
}

func (s *DescribeDDosEventAreaResponse) SetHeaders(v map[string]*string) *DescribeDDosEventAreaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventAreaResponse) SetStatusCode(v int32) *DescribeDDosEventAreaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosEventAreaResponse) SetBody(v *DescribeDDosEventAreaResponseBody) *DescribeDDosEventAreaResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosEventAreaResponse) Validate() error {
	return dara.Validate(s)
}
