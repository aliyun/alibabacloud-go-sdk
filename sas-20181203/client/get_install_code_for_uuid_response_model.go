// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstallCodeForUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstallCodeForUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstallCodeForUuidResponse
	GetStatusCode() *int32
	SetBody(v *GetInstallCodeForUuidResponseBody) *GetInstallCodeForUuidResponse
	GetBody() *GetInstallCodeForUuidResponseBody
}

type GetInstallCodeForUuidResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstallCodeForUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstallCodeForUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstallCodeForUuidResponse) GoString() string {
	return s.String()
}

func (s *GetInstallCodeForUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstallCodeForUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstallCodeForUuidResponse) GetBody() *GetInstallCodeForUuidResponseBody {
	return s.Body
}

func (s *GetInstallCodeForUuidResponse) SetHeaders(v map[string]*string) *GetInstallCodeForUuidResponse {
	s.Headers = v
	return s
}

func (s *GetInstallCodeForUuidResponse) SetStatusCode(v int32) *GetInstallCodeForUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstallCodeForUuidResponse) SetBody(v *GetInstallCodeForUuidResponseBody) *GetInstallCodeForUuidResponse {
	s.Body = v
	return s
}

func (s *GetInstallCodeForUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
