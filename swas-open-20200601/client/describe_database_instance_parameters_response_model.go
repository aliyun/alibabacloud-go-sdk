// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabaseInstanceParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabaseInstanceParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabaseInstanceParametersResponseBody) *DescribeDatabaseInstanceParametersResponse
	GetBody() *DescribeDatabaseInstanceParametersResponseBody
}

type DescribeDatabaseInstanceParametersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstanceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabaseInstanceParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabaseInstanceParametersResponse) GetBody() *DescribeDatabaseInstanceParametersResponseBody {
	return s.Body
}

func (s *DescribeDatabaseInstanceParametersResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetBody(v *DescribeDatabaseInstanceParametersResponseBody) *DescribeDatabaseInstanceParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
