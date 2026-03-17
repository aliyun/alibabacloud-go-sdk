// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagHaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagHaResponseBody) *DescribeSagHaResponse
	GetBody() *DescribeSagHaResponseBody
}

type DescribeSagHaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagHaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagHaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagHaResponse) GetBody() *DescribeSagHaResponseBody {
	return s.Body
}

func (s *DescribeSagHaResponse) SetHeaders(v map[string]*string) *DescribeSagHaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagHaResponse) SetStatusCode(v int32) *DescribeSagHaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagHaResponse) SetBody(v *DescribeSagHaResponseBody) *DescribeSagHaResponse {
	s.Body = v
	return s
}

func (s *DescribeSagHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
