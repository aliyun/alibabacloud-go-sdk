// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMOTokenUsageSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMOTokenUsageSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMOTokenUsageSummaryResponseBody) *DescribeMOTokenUsageSummaryResponse
	GetBody() *DescribeMOTokenUsageSummaryResponseBody
}

type DescribeMOTokenUsageSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMOTokenUsageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMOTokenUsageSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMOTokenUsageSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMOTokenUsageSummaryResponse) GetBody() *DescribeMOTokenUsageSummaryResponseBody {
	return s.Body
}

func (s *DescribeMOTokenUsageSummaryResponse) SetHeaders(v map[string]*string) *DescribeMOTokenUsageSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponse) SetStatusCode(v int32) *DescribeMOTokenUsageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponse) SetBody(v *DescribeMOTokenUsageSummaryResponseBody) *DescribeMOTokenUsageSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
