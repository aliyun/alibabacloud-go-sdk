// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDBNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDBNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckDBNameResponseBody) *CheckDBNameResponse
	GetBody() *CheckDBNameResponseBody
}

type CheckDBNameResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDBNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDBNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameResponse) GoString() string {
	return s.String()
}

func (s *CheckDBNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDBNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDBNameResponse) GetBody() *CheckDBNameResponseBody {
	return s.Body
}

func (s *CheckDBNameResponse) SetHeaders(v map[string]*string) *CheckDBNameResponse {
	s.Headers = v
	return s
}

func (s *CheckDBNameResponse) SetStatusCode(v int32) *CheckDBNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDBNameResponse) SetBody(v *CheckDBNameResponseBody) *CheckDBNameResponse {
	s.Body = v
	return s
}

func (s *CheckDBNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
