// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSoftwaresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallSoftwaresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallSoftwaresResponse
	GetStatusCode() *int32
	SetBody(v *InstallSoftwaresResponseBody) *InstallSoftwaresResponse
	GetBody() *InstallSoftwaresResponseBody
}

type InstallSoftwaresResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallSoftwaresResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallSoftwaresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallSoftwaresResponse) GetBody() *InstallSoftwaresResponseBody {
	return s.Body
}

func (s *InstallSoftwaresResponse) SetHeaders(v map[string]*string) *InstallSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *InstallSoftwaresResponse) SetStatusCode(v int32) *InstallSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallSoftwaresResponse) SetBody(v *InstallSoftwaresResponseBody) *InstallSoftwaresResponse {
	s.Body = v
	return s
}

func (s *InstallSoftwaresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
