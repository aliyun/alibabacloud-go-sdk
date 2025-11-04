// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterImsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterImsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterImsSummaryResponseBody) *DescribeMeterImsSummaryResponse
	GetBody() *DescribeMeterImsSummaryResponseBody
}

type DescribeMeterImsSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterImsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterImsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterImsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterImsSummaryResponse) GetBody() *DescribeMeterImsSummaryResponseBody {
	return s.Body
}

func (s *DescribeMeterImsSummaryResponse) SetHeaders(v map[string]*string) *DescribeMeterImsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImsSummaryResponse) SetStatusCode(v int32) *DescribeMeterImsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImsSummaryResponse) SetBody(v *DescribeMeterImsSummaryResponseBody) *DescribeMeterImsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterImsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
