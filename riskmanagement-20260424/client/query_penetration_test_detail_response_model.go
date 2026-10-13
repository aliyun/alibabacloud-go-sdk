// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPenetrationTestDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPenetrationTestDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryPenetrationTestDetailResponseBody) *QueryPenetrationTestDetailResponse
	GetBody() *QueryPenetrationTestDetailResponseBody
}

type QueryPenetrationTestDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPenetrationTestDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPenetrationTestDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPenetrationTestDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPenetrationTestDetailResponse) GetBody() *QueryPenetrationTestDetailResponseBody {
	return s.Body
}

func (s *QueryPenetrationTestDetailResponse) SetHeaders(v map[string]*string) *QueryPenetrationTestDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryPenetrationTestDetailResponse) SetStatusCode(v int32) *QueryPenetrationTestDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPenetrationTestDetailResponse) SetBody(v *QueryPenetrationTestDetailResponseBody) *QueryPenetrationTestDetailResponse {
	s.Body = v
	return s
}

func (s *QueryPenetrationTestDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
