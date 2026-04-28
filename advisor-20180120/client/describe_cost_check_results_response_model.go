// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostCheckResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostCheckResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostCheckResultsResponseBody) *DescribeCostCheckResultsResponse
	GetBody() *DescribeCostCheckResultsResponseBody
}

type DescribeCostCheckResultsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostCheckResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostCheckResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostCheckResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostCheckResultsResponse) GetBody() *DescribeCostCheckResultsResponseBody {
	return s.Body
}

func (s *DescribeCostCheckResultsResponse) SetHeaders(v map[string]*string) *DescribeCostCheckResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostCheckResultsResponse) SetStatusCode(v int32) *DescribeCostCheckResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostCheckResultsResponse) SetBody(v *DescribeCostCheckResultsResponseBody) *DescribeCostCheckResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeCostCheckResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
