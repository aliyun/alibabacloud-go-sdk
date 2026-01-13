// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiCallDailyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiCallDailyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiCallDailyDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListApiCallDailyDetailResult) *ListApiCallDailyDetailResponse
	GetBody() *ListApiCallDailyDetailResult
}

type ListApiCallDailyDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiCallDailyDetailResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiCallDailyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiCallDailyDetailResponse) GoString() string {
	return s.String()
}

func (s *ListApiCallDailyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiCallDailyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiCallDailyDetailResponse) GetBody() *ListApiCallDailyDetailResult {
	return s.Body
}

func (s *ListApiCallDailyDetailResponse) SetHeaders(v map[string]*string) *ListApiCallDailyDetailResponse {
	s.Headers = v
	return s
}

func (s *ListApiCallDailyDetailResponse) SetStatusCode(v int32) *ListApiCallDailyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiCallDailyDetailResponse) SetBody(v *ListApiCallDailyDetailResult) *ListApiCallDailyDetailResponse {
	s.Body = v
	return s
}

func (s *ListApiCallDailyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
