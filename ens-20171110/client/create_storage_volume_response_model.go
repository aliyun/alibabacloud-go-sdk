// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStorageVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStorageVolumeResponse
	GetStatusCode() *int32
	SetBody(v *CreateStorageVolumeResponseBody) *CreateStorageVolumeResponse
	GetBody() *CreateStorageVolumeResponseBody
}

type CreateStorageVolumeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStorageVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStorageVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageVolumeResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStorageVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStorageVolumeResponse) GetBody() *CreateStorageVolumeResponseBody {
	return s.Body
}

func (s *CreateStorageVolumeResponse) SetHeaders(v map[string]*string) *CreateStorageVolumeResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageVolumeResponse) SetStatusCode(v int32) *CreateStorageVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageVolumeResponse) SetBody(v *CreateStorageVolumeResponseBody) *CreateStorageVolumeResponse {
	s.Body = v
	return s
}

func (s *CreateStorageVolumeResponse) Validate() error {
	return dara.Validate(s)
}
