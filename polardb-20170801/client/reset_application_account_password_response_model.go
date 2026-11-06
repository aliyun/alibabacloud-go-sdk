// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetApplicationAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetApplicationAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetApplicationAccountPasswordResponseBody) *ResetApplicationAccountPasswordResponse
	GetBody() *ResetApplicationAccountPasswordResponseBody
}

type ResetApplicationAccountPasswordResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetApplicationAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetApplicationAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetApplicationAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetApplicationAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetApplicationAccountPasswordResponse) GetBody() *ResetApplicationAccountPasswordResponseBody {
	return s.Body
}

func (s *ResetApplicationAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetApplicationAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetApplicationAccountPasswordResponse) SetStatusCode(v int32) *ResetApplicationAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetApplicationAccountPasswordResponse) SetBody(v *ResetApplicationAccountPasswordResponseBody) *ResetApplicationAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetApplicationAccountPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
