// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpResponse
	GetStatusCode() *int32
	SetBody(v *AddIpResponseBody) *AddIpResponse
	GetBody() *AddIpResponseBody
}

type AddIpResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpResponse) GetBody() *AddIpResponseBody {
	return s.Body
}

func (s *AddIpResponse) SetHeaders(v map[string]*string) *AddIpResponse {
	s.Headers = v
	return s
}

func (s *AddIpResponse) SetStatusCode(v int32) *AddIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

func (s *AddIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
