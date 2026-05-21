// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonthlySlaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMonthlySlaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMonthlySlaListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMonthlySlaListResponseBody) *QueryMonthlySlaListResponse
	GetBody() *QueryMonthlySlaListResponseBody
}

type QueryMonthlySlaListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMonthlySlaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMonthlySlaListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthlySlaListResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthlySlaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMonthlySlaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMonthlySlaListResponse) GetBody() *QueryMonthlySlaListResponseBody {
	return s.Body
}

func (s *QueryMonthlySlaListResponse) SetHeaders(v map[string]*string) *QueryMonthlySlaListResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthlySlaListResponse) SetStatusCode(v int32) *QueryMonthlySlaListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonthlySlaListResponse) SetBody(v *QueryMonthlySlaListResponseBody) *QueryMonthlySlaListResponse {
	s.Body = v
	return s
}

func (s *QueryMonthlySlaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
