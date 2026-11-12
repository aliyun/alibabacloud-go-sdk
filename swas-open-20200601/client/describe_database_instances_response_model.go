// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabaseInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabaseInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabaseInstancesResponseBody) *DescribeDatabaseInstancesResponse
	GetBody() *DescribeDatabaseInstancesResponseBody
}

type DescribeDatabaseInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabaseInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabaseInstancesResponse) GetBody() *DescribeDatabaseInstancesResponseBody {
	return s.Body
}

func (s *DescribeDatabaseInstancesResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetStatusCode(v int32) *DescribeDatabaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetBody(v *DescribeDatabaseInstancesResponseBody) *DescribeDatabaseInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabaseInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
