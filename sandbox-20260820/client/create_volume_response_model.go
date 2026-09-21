// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVolumeResponse
	GetStatusCode() *int32
	SetBody(v *CreateVolumeResponseBody) *CreateVolumeResponse
	GetBody() *CreateVolumeResponseBody
}

type CreateVolumeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeResponse) GoString() string {
	return s.String()
}

func (s *CreateVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVolumeResponse) GetBody() *CreateVolumeResponseBody {
	return s.Body
}

func (s *CreateVolumeResponse) SetHeaders(v map[string]*string) *CreateVolumeResponse {
	s.Headers = v
	return s
}

func (s *CreateVolumeResponse) SetStatusCode(v int32) *CreateVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVolumeResponse) SetBody(v *CreateVolumeResponseBody) *CreateVolumeResponse {
	s.Body = v
	return s
}

func (s *CreateVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
