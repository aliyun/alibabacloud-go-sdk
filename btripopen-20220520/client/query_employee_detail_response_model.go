// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmployeeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEmployeeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEmployeeDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryEmployeeDetailResponseBody) *QueryEmployeeDetailResponse
	GetBody() *QueryEmployeeDetailResponseBody
}

type QueryEmployeeDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmployeeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmployeeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEmployeeDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryEmployeeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEmployeeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEmployeeDetailResponse) GetBody() *QueryEmployeeDetailResponseBody {
	return s.Body
}

func (s *QueryEmployeeDetailResponse) SetHeaders(v map[string]*string) *QueryEmployeeDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryEmployeeDetailResponse) SetStatusCode(v int32) *QueryEmployeeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmployeeDetailResponse) SetBody(v *QueryEmployeeDetailResponseBody) *QueryEmployeeDetailResponse {
	s.Body = v
	return s
}

func (s *QueryEmployeeDetailResponse) Validate() error {
	return dara.Validate(s)
}
