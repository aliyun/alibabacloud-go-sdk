// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnatEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnatEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnatEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnatEntriesResponseBody) *DescribeDnatEntriesResponse
	GetBody() *DescribeDnatEntriesResponseBody
}

type DescribeDnatEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnatEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnatEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnatEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnatEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnatEntriesResponse) GetBody() *DescribeDnatEntriesResponseBody {
	return s.Body
}

func (s *DescribeDnatEntriesResponse) SetHeaders(v map[string]*string) *DescribeDnatEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnatEntriesResponse) SetStatusCode(v int32) *DescribeDnatEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnatEntriesResponse) SetBody(v *DescribeDnatEntriesResponseBody) *DescribeDnatEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeDnatEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
