// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDiskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDiskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDiskAttributeResponseBody) *UpdateDiskAttributeResponse
	GetBody() *UpdateDiskAttributeResponseBody
}

type UpdateDiskAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDiskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDiskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDiskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDiskAttributeResponse) GetBody() *UpdateDiskAttributeResponseBody {
	return s.Body
}

func (s *UpdateDiskAttributeResponse) SetHeaders(v map[string]*string) *UpdateDiskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDiskAttributeResponse) SetStatusCode(v int32) *UpdateDiskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDiskAttributeResponse) SetBody(v *UpdateDiskAttributeResponseBody) *UpdateDiskAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateDiskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
