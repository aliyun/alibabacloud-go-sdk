// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenErrorTimesInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TopTenErrorTimesInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TopTenErrorTimesInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TopTenErrorTimesInstanceResponseBody) *TopTenErrorTimesInstanceResponse
	GetBody() *TopTenErrorTimesInstanceResponseBody
}

type TopTenErrorTimesInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TopTenErrorTimesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TopTenErrorTimesInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TopTenErrorTimesInstanceResponse) GoString() string {
	return s.String()
}

func (s *TopTenErrorTimesInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TopTenErrorTimesInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TopTenErrorTimesInstanceResponse) GetBody() *TopTenErrorTimesInstanceResponseBody {
	return s.Body
}

func (s *TopTenErrorTimesInstanceResponse) SetHeaders(v map[string]*string) *TopTenErrorTimesInstanceResponse {
	s.Headers = v
	return s
}

func (s *TopTenErrorTimesInstanceResponse) SetStatusCode(v int32) *TopTenErrorTimesInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponse) SetBody(v *TopTenErrorTimesInstanceResponseBody) *TopTenErrorTimesInstanceResponse {
	s.Body = v
	return s
}

func (s *TopTenErrorTimesInstanceResponse) Validate() error {
	return dara.Validate(s)
}
