// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceInspectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceInspectionResponseBody) *UpdateInstanceInspectionResponse
	GetBody() *UpdateInstanceInspectionResponseBody
}

type UpdateInstanceInspectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceInspectionResponse) GetBody() *UpdateInstanceInspectionResponseBody {
	return s.Body
}

func (s *UpdateInstanceInspectionResponse) SetHeaders(v map[string]*string) *UpdateInstanceInspectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceInspectionResponse) SetStatusCode(v int32) *UpdateInstanceInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponse) SetBody(v *UpdateInstanceInspectionResponseBody) *UpdateInstanceInspectionResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
