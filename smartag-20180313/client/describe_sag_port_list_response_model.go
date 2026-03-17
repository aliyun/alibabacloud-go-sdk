// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagPortListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagPortListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagPortListResponseBody) *DescribeSagPortListResponse
	GetBody() *DescribeSagPortListResponseBody
}

type DescribeSagPortListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagPortListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagPortListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagPortListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagPortListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagPortListResponse) GetBody() *DescribeSagPortListResponseBody {
	return s.Body
}

func (s *DescribeSagPortListResponse) SetHeaders(v map[string]*string) *DescribeSagPortListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagPortListResponse) SetStatusCode(v int32) *DescribeSagPortListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagPortListResponse) SetBody(v *DescribeSagPortListResponseBody) *DescribeSagPortListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagPortListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
