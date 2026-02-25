// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterSetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateParameterSetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateParameterSetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateParameterSetAttributeResponseBody) *UpdateParameterSetAttributeResponse
	GetBody() *UpdateParameterSetAttributeResponseBody
}

type UpdateParameterSetAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParameterSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParameterSetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateParameterSetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateParameterSetAttributeResponse) GetBody() *UpdateParameterSetAttributeResponseBody {
	return s.Body
}

func (s *UpdateParameterSetAttributeResponse) SetHeaders(v map[string]*string) *UpdateParameterSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateParameterSetAttributeResponse) SetStatusCode(v int32) *UpdateParameterSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParameterSetAttributeResponse) SetBody(v *UpdateParameterSetAttributeResponseBody) *UpdateParameterSetAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateParameterSetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
