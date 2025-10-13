// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAddonReleaseResponseBody) *UpdateAddonReleaseResponse
	GetBody() *UpdateAddonReleaseResponseBody
}

type UpdateAddonReleaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAddonReleaseResponse) GetBody() *UpdateAddonReleaseResponseBody {
	return s.Body
}

func (s *UpdateAddonReleaseResponse) SetHeaders(v map[string]*string) *UpdateAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateAddonReleaseResponse) SetStatusCode(v int32) *UpdateAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAddonReleaseResponse) SetBody(v *UpdateAddonReleaseResponseBody) *UpdateAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *UpdateAddonReleaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
