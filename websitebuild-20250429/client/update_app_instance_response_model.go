// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppInstanceResponseBody) *UpdateAppInstanceResponse
	GetBody() *UpdateAppInstanceResponseBody
}

type UpdateAppInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppInstanceResponse) GetBody() *UpdateAppInstanceResponseBody {
	return s.Body
}

func (s *UpdateAppInstanceResponse) SetHeaders(v map[string]*string) *UpdateAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInstanceResponse) SetStatusCode(v int32) *UpdateAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInstanceResponse) SetBody(v *UpdateAppInstanceResponseBody) *UpdateAppInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
