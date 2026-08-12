// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVolumeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVolumeResponseBody) *UpdateVolumeResponse
	GetBody() *UpdateVolumeResponseBody
}

type UpdateVolumeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVolumeResponse) GetBody() *UpdateVolumeResponseBody {
	return s.Body
}

func (s *UpdateVolumeResponse) SetHeaders(v map[string]*string) *UpdateVolumeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVolumeResponse) SetStatusCode(v int32) *UpdateVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVolumeResponse) SetBody(v *UpdateVolumeResponseBody) *UpdateVolumeResponse {
	s.Body = v
	return s
}

func (s *UpdateVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
