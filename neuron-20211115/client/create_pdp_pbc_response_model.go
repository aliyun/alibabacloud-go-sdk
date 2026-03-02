// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdpPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdpPbcResponse
	GetStatusCode() *int32
	SetBody(v *PdpPbc) *CreatePdpPbcResponse
	GetBody() *PdpPbc
}

type CreatePdpPbcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpPbc            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpPbcResponse) GoString() string {
	return s.String()
}

func (s *CreatePdpPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdpPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdpPbcResponse) GetBody() *PdpPbc {
	return s.Body
}

func (s *CreatePdpPbcResponse) SetHeaders(v map[string]*string) *CreatePdpPbcResponse {
	s.Headers = v
	return s
}

func (s *CreatePdpPbcResponse) SetStatusCode(v int32) *CreatePdpPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdpPbcResponse) SetBody(v *PdpPbc) *CreatePdpPbcResponse {
	s.Body = v
	return s
}

func (s *CreatePdpPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
