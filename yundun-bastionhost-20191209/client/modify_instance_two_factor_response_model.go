// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTwoFactorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceTwoFactorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceTwoFactorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceTwoFactorResponseBody) *ModifyInstanceTwoFactorResponse
	GetBody() *ModifyInstanceTwoFactorResponseBody
}

type ModifyInstanceTwoFactorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceTwoFactorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTwoFactorResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceTwoFactorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceTwoFactorResponse) GetBody() *ModifyInstanceTwoFactorResponseBody {
	return s.Body
}

func (s *ModifyInstanceTwoFactorResponse) SetHeaders(v map[string]*string) *ModifyInstanceTwoFactorResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTwoFactorResponse) SetStatusCode(v int32) *ModifyInstanceTwoFactorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceTwoFactorResponse) SetBody(v *ModifyInstanceTwoFactorResponseBody) *ModifyInstanceTwoFactorResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceTwoFactorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
