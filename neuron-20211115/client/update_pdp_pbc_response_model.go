// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePdpPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePdpPbcResponse
	GetStatusCode() *int32
	SetBody(v *PdpPbc) *UpdatePdpPbcResponse
	GetBody() *PdpPbc
}

type UpdatePdpPbcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpPbc            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpPbcResponse) GoString() string {
	return s.String()
}

func (s *UpdatePdpPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePdpPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePdpPbcResponse) GetBody() *PdpPbc {
	return s.Body
}

func (s *UpdatePdpPbcResponse) SetHeaders(v map[string]*string) *UpdatePdpPbcResponse {
	s.Headers = v
	return s
}

func (s *UpdatePdpPbcResponse) SetStatusCode(v int32) *UpdatePdpPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePdpPbcResponse) SetBody(v *PdpPbc) *UpdatePdpPbcResponse {
	s.Body = v
	return s
}

func (s *UpdatePdpPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
