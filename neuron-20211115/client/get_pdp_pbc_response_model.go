// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpPbcResponse
	GetStatusCode() *int32
	SetBody(v *PdpPbc) *GetPdpPbcResponse
	GetBody() *PdpPbc
}

type GetPdpPbcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpPbc            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpPbcResponse) GoString() string {
	return s.String()
}

func (s *GetPdpPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpPbcResponse) GetBody() *PdpPbc {
	return s.Body
}

func (s *GetPdpPbcResponse) SetHeaders(v map[string]*string) *GetPdpPbcResponse {
	s.Headers = v
	return s
}

func (s *GetPdpPbcResponse) SetStatusCode(v int32) *GetPdpPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpPbcResponse) SetBody(v *PdpPbc) *GetPdpPbcResponse {
	s.Body = v
	return s
}

func (s *GetPdpPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
