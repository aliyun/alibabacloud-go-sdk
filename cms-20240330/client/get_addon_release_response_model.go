// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *GetAddonReleaseResponseBody) *GetAddonReleaseResponse
	GetBody() *GetAddonReleaseResponseBody
}

type GetAddonReleaseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *GetAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAddonReleaseResponse) GetBody() *GetAddonReleaseResponseBody {
	return s.Body
}

func (s *GetAddonReleaseResponse) SetHeaders(v map[string]*string) *GetAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *GetAddonReleaseResponse) SetStatusCode(v int32) *GetAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddonReleaseResponse) SetBody(v *GetAddonReleaseResponseBody) *GetAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *GetAddonReleaseResponse) Validate() error {
	return dara.Validate(s)
}
