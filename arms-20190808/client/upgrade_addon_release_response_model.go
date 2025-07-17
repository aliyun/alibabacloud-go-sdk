// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAddonReleaseResponseBody) *UpgradeAddonReleaseResponse
	GetBody() *UpgradeAddonReleaseResponseBody
}

type UpgradeAddonReleaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAddonReleaseResponse) GetBody() *UpgradeAddonReleaseResponseBody {
	return s.Body
}

func (s *UpgradeAddonReleaseResponse) SetHeaders(v map[string]*string) *UpgradeAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAddonReleaseResponse) SetStatusCode(v int32) *UpgradeAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAddonReleaseResponse) SetBody(v *UpgradeAddonReleaseResponseBody) *UpgradeAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *UpgradeAddonReleaseResponse) Validate() error {
	return dara.Validate(s)
}
