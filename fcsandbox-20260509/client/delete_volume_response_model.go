// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVolumeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVolumeResponseBody) *DeleteVolumeResponse
	GetBody() *DeleteVolumeResponseBody
}

type DeleteVolumeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVolumeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVolumeResponse) GetBody() *DeleteVolumeResponseBody {
	return s.Body
}

func (s *DeleteVolumeResponse) SetHeaders(v map[string]*string) *DeleteVolumeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVolumeResponse) SetStatusCode(v int32) *DeleteVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVolumeResponse) SetBody(v *DeleteVolumeResponseBody) *DeleteVolumeResponse {
	s.Body = v
	return s
}

func (s *DeleteVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
