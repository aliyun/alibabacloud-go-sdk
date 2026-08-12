// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVolumeResponse
	GetStatusCode() *int32
	SetBody(v *GetVolumeResponseBody) *GetVolumeResponse
	GetBody() *GetVolumeResponseBody
}

type GetVolumeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVolumeResponse) GoString() string {
	return s.String()
}

func (s *GetVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVolumeResponse) GetBody() *GetVolumeResponseBody {
	return s.Body
}

func (s *GetVolumeResponse) SetHeaders(v map[string]*string) *GetVolumeResponse {
	s.Headers = v
	return s
}

func (s *GetVolumeResponse) SetStatusCode(v int32) *GetVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVolumeResponse) SetBody(v *GetVolumeResponseBody) *GetVolumeResponse {
	s.Body = v
	return s
}

func (s *GetVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
