// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataSourceConnectivityByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataSourceConnectivityByIdResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataSourceConnectivityByIdResponseBody) *CheckDataSourceConnectivityByIdResponse
	GetBody() *CheckDataSourceConnectivityByIdResponseBody
}

type CheckDataSourceConnectivityByIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataSourceConnectivityByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataSourceConnectivityByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdResponse) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataSourceConnectivityByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataSourceConnectivityByIdResponse) GetBody() *CheckDataSourceConnectivityByIdResponseBody {
	return s.Body
}

func (s *CheckDataSourceConnectivityByIdResponse) SetHeaders(v map[string]*string) *CheckDataSourceConnectivityByIdResponse {
	s.Headers = v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponse) SetStatusCode(v int32) *CheckDataSourceConnectivityByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponse) SetBody(v *CheckDataSourceConnectivityByIdResponseBody) *CheckDataSourceConnectivityByIdResponse {
	s.Body = v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponse) Validate() error {
	return dara.Validate(s)
}
