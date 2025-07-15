// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTimeShiftDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainTimeShiftDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainTimeShiftDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainTimeShiftDataResponseBody) *DescribeLiveDomainTimeShiftDataResponse
	GetBody() *DescribeLiveDomainTimeShiftDataResponseBody
}

type DescribeLiveDomainTimeShiftDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainTimeShiftDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainTimeShiftDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainTimeShiftDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainTimeShiftDataResponse) GetBody() *DescribeLiveDomainTimeShiftDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainTimeShiftDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainTimeShiftDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponse) SetStatusCode(v int32) *DescribeLiveDomainTimeShiftDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponse) SetBody(v *DescribeLiveDomainTimeShiftDataResponseBody) *DescribeLiveDomainTimeShiftDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponse) Validate() error {
	return dara.Validate(s)
}
