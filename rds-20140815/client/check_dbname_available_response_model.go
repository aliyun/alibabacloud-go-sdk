// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameAvailableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDBNameAvailableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDBNameAvailableResponse
	GetStatusCode() *int32
	SetBody(v *CheckDBNameAvailableResponseBody) *CheckDBNameAvailableResponse
	GetBody() *CheckDBNameAvailableResponseBody
}

type CheckDBNameAvailableResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDBNameAvailableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDBNameAvailableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameAvailableResponse) GoString() string {
	return s.String()
}

func (s *CheckDBNameAvailableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDBNameAvailableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDBNameAvailableResponse) GetBody() *CheckDBNameAvailableResponseBody {
	return s.Body
}

func (s *CheckDBNameAvailableResponse) SetHeaders(v map[string]*string) *CheckDBNameAvailableResponse {
	s.Headers = v
	return s
}

func (s *CheckDBNameAvailableResponse) SetStatusCode(v int32) *CheckDBNameAvailableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDBNameAvailableResponse) SetBody(v *CheckDBNameAvailableResponseBody) *CheckDBNameAvailableResponse {
	s.Body = v
	return s
}

func (s *CheckDBNameAvailableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
