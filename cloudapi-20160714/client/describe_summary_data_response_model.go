// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSummaryDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSummaryDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSummaryDataResponseBody) *DescribeSummaryDataResponse
	GetBody() *DescribeSummaryDataResponseBody
}

type DescribeSummaryDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSummaryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSummaryDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSummaryDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSummaryDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSummaryDataResponse) GetBody() *DescribeSummaryDataResponseBody {
	return s.Body
}

func (s *DescribeSummaryDataResponse) SetHeaders(v map[string]*string) *DescribeSummaryDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSummaryDataResponse) SetStatusCode(v int32) *DescribeSummaryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSummaryDataResponse) SetBody(v *DescribeSummaryDataResponseBody) *DescribeSummaryDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSummaryDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
