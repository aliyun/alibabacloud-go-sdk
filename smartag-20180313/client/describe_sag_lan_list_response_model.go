// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagLanListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagLanListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagLanListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagLanListResponseBody) *DescribeSagLanListResponse
	GetBody() *DescribeSagLanListResponseBody
}

type DescribeSagLanListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagLanListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagLanListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagLanListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagLanListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagLanListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagLanListResponse) GetBody() *DescribeSagLanListResponseBody {
	return s.Body
}

func (s *DescribeSagLanListResponse) SetHeaders(v map[string]*string) *DescribeSagLanListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagLanListResponse) SetStatusCode(v int32) *DescribeSagLanListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagLanListResponse) SetBody(v *DescribeSagLanListResponseBody) *DescribeSagLanListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagLanListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
