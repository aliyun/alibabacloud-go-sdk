// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySuccessIcpDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySuccessIcpDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySuccessIcpDataResponse
	GetStatusCode() *int32
	SetBody(v *QuerySuccessIcpDataResponseBody) *QuerySuccessIcpDataResponse
	GetBody() *QuerySuccessIcpDataResponseBody
}

type QuerySuccessIcpDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySuccessIcpDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySuccessIcpDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySuccessIcpDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySuccessIcpDataResponse) GetBody() *QuerySuccessIcpDataResponseBody {
	return s.Body
}

func (s *QuerySuccessIcpDataResponse) SetHeaders(v map[string]*string) *QuerySuccessIcpDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySuccessIcpDataResponse) SetStatusCode(v int32) *QuerySuccessIcpDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySuccessIcpDataResponse) SetBody(v *QuerySuccessIcpDataResponseBody) *QuerySuccessIcpDataResponse {
	s.Body = v
	return s
}

func (s *QuerySuccessIcpDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
