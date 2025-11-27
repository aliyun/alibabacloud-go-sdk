// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceExistResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceExistResponseBody) *CheckInstanceExistResponse
	GetBody() *CheckInstanceExistResponseBody
}

type CheckInstanceExistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceExistResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceExistResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceExistResponse) GetBody() *CheckInstanceExistResponseBody {
	return s.Body
}

func (s *CheckInstanceExistResponse) SetHeaders(v map[string]*string) *CheckInstanceExistResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceExistResponse) SetStatusCode(v int32) *CheckInstanceExistResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceExistResponse) SetBody(v *CheckInstanceExistResponseBody) *CheckInstanceExistResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceExistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
