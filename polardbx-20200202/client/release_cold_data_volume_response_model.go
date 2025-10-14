// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseColdDataVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseColdDataVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseColdDataVolumeResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseColdDataVolumeResponseBody) *ReleaseColdDataVolumeResponse
	GetBody() *ReleaseColdDataVolumeResponseBody
}

type ReleaseColdDataVolumeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseColdDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseColdDataVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseColdDataVolumeResponse) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseColdDataVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseColdDataVolumeResponse) GetBody() *ReleaseColdDataVolumeResponseBody {
	return s.Body
}

func (s *ReleaseColdDataVolumeResponse) SetHeaders(v map[string]*string) *ReleaseColdDataVolumeResponse {
	s.Headers = v
	return s
}

func (s *ReleaseColdDataVolumeResponse) SetStatusCode(v int32) *ReleaseColdDataVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseColdDataVolumeResponse) SetBody(v *ReleaseColdDataVolumeResponseBody) *ReleaseColdDataVolumeResponse {
	s.Body = v
	return s
}

func (s *ReleaseColdDataVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
