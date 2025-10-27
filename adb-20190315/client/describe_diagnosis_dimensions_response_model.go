// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisDimensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse
	GetBody() *DescribeDiagnosisDimensionsResponseBody
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisDimensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisDimensionsResponse) GetBody() *DescribeDiagnosisDimensionsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
