// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabasesZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabasesZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabasesZonalResponseBody) *DescribeDatabasesZonalResponse
	GetBody() *DescribeDatabasesZonalResponseBody
}

type DescribeDatabasesZonalResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabasesZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabasesZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabasesZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabasesZonalResponse) GetBody() *DescribeDatabasesZonalResponseBody {
	return s.Body
}

func (s *DescribeDatabasesZonalResponse) SetHeaders(v map[string]*string) *DescribeDatabasesZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesZonalResponse) SetStatusCode(v int32) *DescribeDatabasesZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabasesZonalResponse) SetBody(v *DescribeDatabasesZonalResponseBody) *DescribeDatabasesZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabasesZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
