// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineServiceVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnlineServiceVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnlineServiceVolumeResponse
	GetStatusCode() *int32
	SetBody(v *GetOnlineServiceVolumeResponseBody) *GetOnlineServiceVolumeResponse
	GetBody() *GetOnlineServiceVolumeResponseBody
}

type GetOnlineServiceVolumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnlineServiceVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnlineServiceVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineServiceVolumeResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnlineServiceVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnlineServiceVolumeResponse) GetBody() *GetOnlineServiceVolumeResponseBody {
	return s.Body
}

func (s *GetOnlineServiceVolumeResponse) SetHeaders(v map[string]*string) *GetOnlineServiceVolumeResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineServiceVolumeResponse) SetStatusCode(v int32) *GetOnlineServiceVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnlineServiceVolumeResponse) SetBody(v *GetOnlineServiceVolumeResponseBody) *GetOnlineServiceVolumeResponse {
	s.Body = v
	return s
}

func (s *GetOnlineServiceVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
