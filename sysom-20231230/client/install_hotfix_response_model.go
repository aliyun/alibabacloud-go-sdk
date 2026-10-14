// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHotfixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallHotfixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallHotfixResponse
	GetStatusCode() *int32
	SetBody(v *InstallHotfixResponseBody) *InstallHotfixResponse
	GetBody() *InstallHotfixResponseBody
}

type InstallHotfixResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallHotfixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallHotfixResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallHotfixResponse) GoString() string {
	return s.String()
}

func (s *InstallHotfixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallHotfixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallHotfixResponse) GetBody() *InstallHotfixResponseBody {
	return s.Body
}

func (s *InstallHotfixResponse) SetHeaders(v map[string]*string) *InstallHotfixResponse {
	s.Headers = v
	return s
}

func (s *InstallHotfixResponse) SetStatusCode(v int32) *InstallHotfixResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallHotfixResponse) SetBody(v *InstallHotfixResponseBody) *InstallHotfixResponse {
	s.Body = v
	return s
}

func (s *InstallHotfixResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
