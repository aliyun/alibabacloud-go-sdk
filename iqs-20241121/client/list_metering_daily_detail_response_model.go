// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeteringDailyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMeteringDailyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMeteringDailyDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListMeteringDailyDetailResult) *ListMeteringDailyDetailResponse
	GetBody() *ListMeteringDailyDetailResult
}

type ListMeteringDailyDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMeteringDailyDetailResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMeteringDailyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMeteringDailyDetailResponse) GoString() string {
	return s.String()
}

func (s *ListMeteringDailyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMeteringDailyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMeteringDailyDetailResponse) GetBody() *ListMeteringDailyDetailResult {
	return s.Body
}

func (s *ListMeteringDailyDetailResponse) SetHeaders(v map[string]*string) *ListMeteringDailyDetailResponse {
	s.Headers = v
	return s
}

func (s *ListMeteringDailyDetailResponse) SetStatusCode(v int32) *ListMeteringDailyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMeteringDailyDetailResponse) SetBody(v *ListMeteringDailyDetailResult) *ListMeteringDailyDetailResponse {
	s.Body = v
	return s
}

func (s *ListMeteringDailyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
