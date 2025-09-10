// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAddonReleaseResponseBody) *DeleteAddonReleaseResponse
	GetBody() *DeleteAddonReleaseResponseBody
}

type DeleteAddonReleaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAddonReleaseResponse) GetBody() *DeleteAddonReleaseResponseBody {
	return s.Body
}

func (s *DeleteAddonReleaseResponse) SetHeaders(v map[string]*string) *DeleteAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddonReleaseResponse) SetStatusCode(v int32) *DeleteAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAddonReleaseResponse) SetBody(v *DeleteAddonReleaseResponseBody) *DeleteAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *DeleteAddonReleaseResponse) Validate() error {
	return dara.Validate(s)
}
