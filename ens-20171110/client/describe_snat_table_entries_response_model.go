// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnatTableEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnatTableEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnatTableEntriesResponseBody) *DescribeSnatTableEntriesResponse
	GetBody() *DescribeSnatTableEntriesResponseBody
}

type DescribeSnatTableEntriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnatTableEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnatTableEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnatTableEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnatTableEntriesResponse) GetBody() *DescribeSnatTableEntriesResponseBody {
	return s.Body
}

func (s *DescribeSnatTableEntriesResponse) SetHeaders(v map[string]*string) *DescribeSnatTableEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnatTableEntriesResponse) SetStatusCode(v int32) *DescribeSnatTableEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnatTableEntriesResponse) SetBody(v *DescribeSnatTableEntriesResponseBody) *DescribeSnatTableEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeSnatTableEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
