// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBruteForceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBruteForceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBruteForceSummaryResponseBody) *DescribeBruteForceSummaryResponse
	GetBody() *DescribeBruteForceSummaryResponseBody
}

type DescribeBruteForceSummaryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBruteForceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBruteForceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBruteForceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBruteForceSummaryResponse) GetBody() *DescribeBruteForceSummaryResponseBody {
	return s.Body
}

func (s *DescribeBruteForceSummaryResponse) SetHeaders(v map[string]*string) *DescribeBruteForceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBruteForceSummaryResponse) SetStatusCode(v int32) *DescribeBruteForceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBruteForceSummaryResponse) SetBody(v *DescribeBruteForceSummaryResponseBody) *DescribeBruteForceSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeBruteForceSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
