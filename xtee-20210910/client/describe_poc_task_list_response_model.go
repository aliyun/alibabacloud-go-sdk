// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePocTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePocTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePocTaskListResponseBody) *DescribePocTaskListResponse
	GetBody() *DescribePocTaskListResponseBody
}

type DescribePocTaskListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePocTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePocTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePocTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePocTaskListResponse) GetBody() *DescribePocTaskListResponseBody {
	return s.Body
}

func (s *DescribePocTaskListResponse) SetHeaders(v map[string]*string) *DescribePocTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribePocTaskListResponse) SetStatusCode(v int32) *DescribePocTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePocTaskListResponse) SetBody(v *DescribePocTaskListResponseBody) *DescribePocTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribePocTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
