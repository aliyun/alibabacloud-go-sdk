// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDriveResponse
	GetStatusCode() *int32
	SetBody(v *Drive) *UpdateDriveResponse
	GetBody() *Drive
}

type UpdateDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDriveResponse) GoString() string {
	return s.String()
}

func (s *UpdateDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDriveResponse) GetBody() *Drive {
	return s.Body
}

func (s *UpdateDriveResponse) SetHeaders(v map[string]*string) *UpdateDriveResponse {
	s.Headers = v
	return s
}

func (s *UpdateDriveResponse) SetStatusCode(v int32) *UpdateDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDriveResponse) SetBody(v *Drive) *UpdateDriveResponse {
	s.Body = v
	return s
}

func (s *UpdateDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
