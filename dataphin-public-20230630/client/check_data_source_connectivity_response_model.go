// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataSourceConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataSourceConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataSourceConnectivityResponseBody) *CheckDataSourceConnectivityResponse
	GetBody() *CheckDataSourceConnectivityResponseBody
}

type CheckDataSourceConnectivityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataSourceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataSourceConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataSourceConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataSourceConnectivityResponse) GetBody() *CheckDataSourceConnectivityResponseBody {
	return s.Body
}

func (s *CheckDataSourceConnectivityResponse) SetHeaders(v map[string]*string) *CheckDataSourceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *CheckDataSourceConnectivityResponse) SetStatusCode(v int32) *CheckDataSourceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityResponse) SetBody(v *CheckDataSourceConnectivityResponseBody) *CheckDataSourceConnectivityResponse {
	s.Body = v
	return s
}

func (s *CheckDataSourceConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
