// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallHotfixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallHotfixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallHotfixResponse
	GetStatusCode() *int32
	SetBody(v *UninstallHotfixResponseBody) *UninstallHotfixResponse
	GetBody() *UninstallHotfixResponseBody
}

type UninstallHotfixResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallHotfixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallHotfixResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallHotfixResponse) GoString() string {
	return s.String()
}

func (s *UninstallHotfixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallHotfixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallHotfixResponse) GetBody() *UninstallHotfixResponseBody {
	return s.Body
}

func (s *UninstallHotfixResponse) SetHeaders(v map[string]*string) *UninstallHotfixResponse {
	s.Headers = v
	return s
}

func (s *UninstallHotfixResponse) SetStatusCode(v int32) *UninstallHotfixResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallHotfixResponse) SetBody(v *UninstallHotfixResponseBody) *UninstallHotfixResponse {
	s.Body = v
	return s
}

func (s *UninstallHotfixResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
