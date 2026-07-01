// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRCSSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRCSSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRCSSignatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRCSSignatureResponseBody) *UpdateRCSSignatureResponse
	GetBody() *UpdateRCSSignatureResponseBody
}

type UpdateRCSSignatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRCSSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRCSSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRCSSignatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateRCSSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRCSSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRCSSignatureResponse) GetBody() *UpdateRCSSignatureResponseBody {
	return s.Body
}

func (s *UpdateRCSSignatureResponse) SetHeaders(v map[string]*string) *UpdateRCSSignatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateRCSSignatureResponse) SetStatusCode(v int32) *UpdateRCSSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRCSSignatureResponse) SetBody(v *UpdateRCSSignatureResponseBody) *UpdateRCSSignatureResponse {
	s.Body = v
	return s
}

func (s *UpdateRCSSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
