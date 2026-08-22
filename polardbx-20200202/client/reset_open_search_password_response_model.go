// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetOpenSearchPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetOpenSearchPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetOpenSearchPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetOpenSearchPasswordResponseBody) *ResetOpenSearchPasswordResponse
	GetBody() *ResetOpenSearchPasswordResponseBody
}

type ResetOpenSearchPasswordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetOpenSearchPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetOpenSearchPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetOpenSearchPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetOpenSearchPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetOpenSearchPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetOpenSearchPasswordResponse) GetBody() *ResetOpenSearchPasswordResponseBody {
	return s.Body
}

func (s *ResetOpenSearchPasswordResponse) SetHeaders(v map[string]*string) *ResetOpenSearchPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetOpenSearchPasswordResponse) SetStatusCode(v int32) *ResetOpenSearchPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetOpenSearchPasswordResponse) SetBody(v *ResetOpenSearchPasswordResponseBody) *ResetOpenSearchPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetOpenSearchPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
