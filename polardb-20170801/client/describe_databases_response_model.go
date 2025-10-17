// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse
	GetBody() *DescribeDatabasesResponseBody
}

type DescribeDatabasesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabasesResponse) GetBody() *DescribeDatabasesResponseBody {
	return s.Body
}

func (s *DescribeDatabasesResponse) SetHeaders(v map[string]*string) *DescribeDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesResponse) SetStatusCode(v int32) *DescribeDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabasesResponse) SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
