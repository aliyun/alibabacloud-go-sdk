// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarClawMCPServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPolarClawMCPServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPolarClawMCPServerResponse
	GetStatusCode() *int32
	SetBody(v *AddPolarClawMCPServerResponseBody) *AddPolarClawMCPServerResponse
	GetBody() *AddPolarClawMCPServerResponseBody
}

type AddPolarClawMCPServerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPolarClawMCPServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPolarClawMCPServerResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPolarClawMCPServerResponse) GoString() string {
	return s.String()
}

func (s *AddPolarClawMCPServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPolarClawMCPServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPolarClawMCPServerResponse) GetBody() *AddPolarClawMCPServerResponseBody {
	return s.Body
}

func (s *AddPolarClawMCPServerResponse) SetHeaders(v map[string]*string) *AddPolarClawMCPServerResponse {
	s.Headers = v
	return s
}

func (s *AddPolarClawMCPServerResponse) SetStatusCode(v int32) *AddPolarClawMCPServerResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPolarClawMCPServerResponse) SetBody(v *AddPolarClawMCPServerResponseBody) *AddPolarClawMCPServerResponse {
	s.Body = v
	return s
}

func (s *AddPolarClawMCPServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
