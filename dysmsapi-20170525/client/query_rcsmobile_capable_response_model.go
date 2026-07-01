// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRCSMobileCapableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRCSMobileCapableResponse
	GetStatusCode() *int32
	SetBody(v *QueryRCSMobileCapableResponseBody) *QueryRCSMobileCapableResponse
	GetBody() *QueryRCSMobileCapableResponseBody
}

type QueryRCSMobileCapableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRCSMobileCapableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRCSMobileCapableResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableResponse) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRCSMobileCapableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRCSMobileCapableResponse) GetBody() *QueryRCSMobileCapableResponseBody {
	return s.Body
}

func (s *QueryRCSMobileCapableResponse) SetHeaders(v map[string]*string) *QueryRCSMobileCapableResponse {
	s.Headers = v
	return s
}

func (s *QueryRCSMobileCapableResponse) SetStatusCode(v int32) *QueryRCSMobileCapableResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRCSMobileCapableResponse) SetBody(v *QueryRCSMobileCapableResponseBody) *QueryRCSMobileCapableResponse {
	s.Body = v
	return s
}

func (s *QueryRCSMobileCapableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
