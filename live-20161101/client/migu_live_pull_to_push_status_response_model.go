// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MiguLivePullToPushStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MiguLivePullToPushStatusResponse
	GetStatusCode() *int32
	SetBody(v *MiguLivePullToPushStatusResponseBody) *MiguLivePullToPushStatusResponse
	GetBody() *MiguLivePullToPushStatusResponseBody
}

type MiguLivePullToPushStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MiguLivePullToPushStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MiguLivePullToPushStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStatusResponse) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MiguLivePullToPushStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MiguLivePullToPushStatusResponse) GetBody() *MiguLivePullToPushStatusResponseBody {
	return s.Body
}

func (s *MiguLivePullToPushStatusResponse) SetHeaders(v map[string]*string) *MiguLivePullToPushStatusResponse {
	s.Headers = v
	return s
}

func (s *MiguLivePullToPushStatusResponse) SetStatusCode(v int32) *MiguLivePullToPushStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *MiguLivePullToPushStatusResponse) SetBody(v *MiguLivePullToPushStatusResponseBody) *MiguLivePullToPushStatusResponse {
	s.Body = v
	return s
}

func (s *MiguLivePullToPushStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
