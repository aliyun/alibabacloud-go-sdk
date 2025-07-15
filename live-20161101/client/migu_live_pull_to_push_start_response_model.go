// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MiguLivePullToPushStartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MiguLivePullToPushStartResponse
	GetStatusCode() *int32
	SetBody(v *MiguLivePullToPushStartResponseBody) *MiguLivePullToPushStartResponse
	GetBody() *MiguLivePullToPushStartResponseBody
}

type MiguLivePullToPushStartResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MiguLivePullToPushStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MiguLivePullToPushStartResponse) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStartResponse) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MiguLivePullToPushStartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MiguLivePullToPushStartResponse) GetBody() *MiguLivePullToPushStartResponseBody {
	return s.Body
}

func (s *MiguLivePullToPushStartResponse) SetHeaders(v map[string]*string) *MiguLivePullToPushStartResponse {
	s.Headers = v
	return s
}

func (s *MiguLivePullToPushStartResponse) SetStatusCode(v int32) *MiguLivePullToPushStartResponse {
	s.StatusCode = &v
	return s
}

func (s *MiguLivePullToPushStartResponse) SetBody(v *MiguLivePullToPushStartResponseBody) *MiguLivePullToPushStartResponse {
	s.Body = v
	return s
}

func (s *MiguLivePullToPushStartResponse) Validate() error {
	return dara.Validate(s)
}
