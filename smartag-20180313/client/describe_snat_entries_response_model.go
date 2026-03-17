// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnatEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnatEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnatEntriesResponseBody) *DescribeSnatEntriesResponse
	GetBody() *DescribeSnatEntriesResponseBody
}

type DescribeSnatEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnatEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnatEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnatEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnatEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnatEntriesResponse) GetBody() *DescribeSnatEntriesResponseBody {
	return s.Body
}

func (s *DescribeSnatEntriesResponse) SetHeaders(v map[string]*string) *DescribeSnatEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnatEntriesResponse) SetStatusCode(v int32) *DescribeSnatEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnatEntriesResponse) SetBody(v *DescribeSnatEntriesResponseBody) *DescribeSnatEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeSnatEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
