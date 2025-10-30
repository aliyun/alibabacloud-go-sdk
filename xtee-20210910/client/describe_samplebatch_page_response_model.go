// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSamplebatchPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSamplebatchPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSamplebatchPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSamplebatchPageResponseBody) *DescribeSamplebatchPageResponse
	GetBody() *DescribeSamplebatchPageResponseBody
}

type DescribeSamplebatchPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSamplebatchPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSamplebatchPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSamplebatchPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeSamplebatchPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSamplebatchPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSamplebatchPageResponse) GetBody() *DescribeSamplebatchPageResponseBody {
	return s.Body
}

func (s *DescribeSamplebatchPageResponse) SetHeaders(v map[string]*string) *DescribeSamplebatchPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeSamplebatchPageResponse) SetStatusCode(v int32) *DescribeSamplebatchPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSamplebatchPageResponse) SetBody(v *DescribeSamplebatchPageResponseBody) *DescribeSamplebatchPageResponse {
	s.Body = v
	return s
}

func (s *DescribeSamplebatchPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
