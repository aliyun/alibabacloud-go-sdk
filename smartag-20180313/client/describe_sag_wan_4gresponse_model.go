// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWan4GResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagWan4GResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagWan4GResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagWan4GResponseBody) *DescribeSagWan4GResponse
	GetBody() *DescribeSagWan4GResponseBody
}

type DescribeSagWan4GResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagWan4GResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagWan4GResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWan4GResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagWan4GResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagWan4GResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagWan4GResponse) GetBody() *DescribeSagWan4GResponseBody {
	return s.Body
}

func (s *DescribeSagWan4GResponse) SetHeaders(v map[string]*string) *DescribeSagWan4GResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagWan4GResponse) SetStatusCode(v int32) *DescribeSagWan4GResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagWan4GResponse) SetBody(v *DescribeSagWan4GResponseBody) *DescribeSagWan4GResponse {
	s.Body = v
	return s
}

func (s *DescribeSagWan4GResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
