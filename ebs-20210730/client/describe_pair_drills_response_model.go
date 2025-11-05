// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePairDrillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePairDrillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePairDrillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePairDrillsResponseBody) *DescribePairDrillsResponse
	GetBody() *DescribePairDrillsResponseBody
}

type DescribePairDrillsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePairDrillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePairDrillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePairDrillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePairDrillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePairDrillsResponse) GetBody() *DescribePairDrillsResponseBody {
	return s.Body
}

func (s *DescribePairDrillsResponse) SetHeaders(v map[string]*string) *DescribePairDrillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePairDrillsResponse) SetStatusCode(v int32) *DescribePairDrillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePairDrillsResponse) SetBody(v *DescribePairDrillsResponseBody) *DescribePairDrillsResponse {
	s.Body = v
	return s
}

func (s *DescribePairDrillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
