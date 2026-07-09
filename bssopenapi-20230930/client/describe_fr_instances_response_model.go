// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFrInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFrInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFrInstancesResponseBody) *DescribeFrInstancesResponse
	GetBody() *DescribeFrInstancesResponseBody
}

type DescribeFrInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFrInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFrInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFrInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFrInstancesResponse) GetBody() *DescribeFrInstancesResponseBody {
	return s.Body
}

func (s *DescribeFrInstancesResponse) SetHeaders(v map[string]*string) *DescribeFrInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFrInstancesResponse) SetStatusCode(v int32) *DescribeFrInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFrInstancesResponse) SetBody(v *DescribeFrInstancesResponseBody) *DescribeFrInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeFrInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
