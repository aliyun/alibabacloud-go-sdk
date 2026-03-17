// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagWanListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagWanListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagWanListResponseBody) *DescribeSagWanListResponse
	GetBody() *DescribeSagWanListResponseBody
}

type DescribeSagWanListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagWanListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagWanListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagWanListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagWanListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagWanListResponse) GetBody() *DescribeSagWanListResponseBody {
	return s.Body
}

func (s *DescribeSagWanListResponse) SetHeaders(v map[string]*string) *DescribeSagWanListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagWanListResponse) SetStatusCode(v int32) *DescribeSagWanListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagWanListResponse) SetBody(v *DescribeSagWanListResponseBody) *DescribeSagWanListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagWanListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
