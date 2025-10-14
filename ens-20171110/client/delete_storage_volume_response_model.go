// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStorageVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStorageVolumeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStorageVolumeResponseBody) *DeleteStorageVolumeResponse
	GetBody() *DeleteStorageVolumeResponseBody
}

type DeleteStorageVolumeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStorageVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStorageVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageVolumeResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStorageVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStorageVolumeResponse) GetBody() *DeleteStorageVolumeResponseBody {
	return s.Body
}

func (s *DeleteStorageVolumeResponse) SetHeaders(v map[string]*string) *DeleteStorageVolumeResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageVolumeResponse) SetStatusCode(v int32) *DeleteStorageVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageVolumeResponse) SetBody(v *DeleteStorageVolumeResponseBody) *DeleteStorageVolumeResponse {
	s.Body = v
	return s
}

func (s *DeleteStorageVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
