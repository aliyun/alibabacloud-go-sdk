// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDsReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDsReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDsReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDsReportsResponseBody) *DescribeDsReportsResponse
	GetBody() *DescribeDsReportsResponseBody
}

type DescribeDsReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDsReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDsReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDsReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDsReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDsReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDsReportsResponse) GetBody() *DescribeDsReportsResponseBody {
	return s.Body
}

func (s *DescribeDsReportsResponse) SetHeaders(v map[string]*string) *DescribeDsReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDsReportsResponse) SetStatusCode(v int32) *DescribeDsReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDsReportsResponse) SetBody(v *DescribeDsReportsResponseBody) *DescribeDsReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeDsReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
