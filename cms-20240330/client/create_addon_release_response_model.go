// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateAddonReleaseResponseBody) *CreateAddonReleaseResponse
	GetBody() *CreateAddonReleaseResponseBody
}

type CreateAddonReleaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *CreateAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAddonReleaseResponse) GetBody() *CreateAddonReleaseResponseBody {
	return s.Body
}

func (s *CreateAddonReleaseResponse) SetHeaders(v map[string]*string) *CreateAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *CreateAddonReleaseResponse) SetStatusCode(v int32) *CreateAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAddonReleaseResponse) SetBody(v *CreateAddonReleaseResponseBody) *CreateAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *CreateAddonReleaseResponse) Validate() error {
	return dara.Validate(s)
}
