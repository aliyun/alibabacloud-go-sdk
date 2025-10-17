// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDSEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDSEntityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDSEntityResponseBody) *UpdateDSEntityResponse
	GetBody() *UpdateDSEntityResponseBody
}

type UpdateDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDSEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDSEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDSEntityResponse) GetBody() *UpdateDSEntityResponseBody {
	return s.Body
}

func (s *UpdateDSEntityResponse) SetHeaders(v map[string]*string) *UpdateDSEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateDSEntityResponse) SetStatusCode(v int32) *UpdateDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDSEntityResponse) SetBody(v *UpdateDSEntityResponseBody) *UpdateDSEntityResponse {
	s.Body = v
	return s
}

func (s *UpdateDSEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
