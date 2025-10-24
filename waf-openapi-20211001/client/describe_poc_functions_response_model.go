// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocFunctionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePocFunctionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePocFunctionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePocFunctionsResponseBody) *DescribePocFunctionsResponse
	GetBody() *DescribePocFunctionsResponseBody
}

type DescribePocFunctionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePocFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePocFunctionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePocFunctionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePocFunctionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePocFunctionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePocFunctionsResponse) GetBody() *DescribePocFunctionsResponseBody {
	return s.Body
}

func (s *DescribePocFunctionsResponse) SetHeaders(v map[string]*string) *DescribePocFunctionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePocFunctionsResponse) SetStatusCode(v int32) *DescribePocFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePocFunctionsResponse) SetBody(v *DescribePocFunctionsResponseBody) *DescribePocFunctionsResponse {
	s.Body = v
	return s
}

func (s *DescribePocFunctionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
