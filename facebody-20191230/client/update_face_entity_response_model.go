// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFaceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFaceEntityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFaceEntityResponseBody) *UpdateFaceEntityResponse
	GetBody() *UpdateFaceEntityResponseBody
}

type UpdateFaceEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFaceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFaceEntityResponse) GetBody() *UpdateFaceEntityResponseBody {
	return s.Body
}

func (s *UpdateFaceEntityResponse) SetHeaders(v map[string]*string) *UpdateFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceEntityResponse) SetStatusCode(v int32) *UpdateFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceEntityResponse) SetBody(v *UpdateFaceEntityResponseBody) *UpdateFaceEntityResponse {
	s.Body = v
	return s
}

func (s *UpdateFaceEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
