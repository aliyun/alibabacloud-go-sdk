// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAegisForLingjunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAegisForLingjunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAegisForLingjunResponse
	GetStatusCode() *int32
	SetBody(v *InstallAegisForLingjunResponseBody) *InstallAegisForLingjunResponse
	GetBody() *InstallAegisForLingjunResponseBody
}

type InstallAegisForLingjunResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAegisForLingjunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAegisForLingjunResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAegisForLingjunResponse) GoString() string {
	return s.String()
}

func (s *InstallAegisForLingjunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAegisForLingjunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAegisForLingjunResponse) GetBody() *InstallAegisForLingjunResponseBody {
	return s.Body
}

func (s *InstallAegisForLingjunResponse) SetHeaders(v map[string]*string) *InstallAegisForLingjunResponse {
	s.Headers = v
	return s
}

func (s *InstallAegisForLingjunResponse) SetStatusCode(v int32) *InstallAegisForLingjunResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAegisForLingjunResponse) SetBody(v *InstallAegisForLingjunResponseBody) *InstallAegisForLingjunResponse {
	s.Body = v
	return s
}

func (s *InstallAegisForLingjunResponse) Validate() error {
	return dara.Validate(s)
}
