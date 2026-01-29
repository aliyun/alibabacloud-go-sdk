// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeResourcePackageResponseBody) *UpgradeResourcePackageResponse
	GetBody() *UpgradeResourcePackageResponseBody
}

type UpgradeResourcePackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeResourcePackageResponse) GetBody() *UpgradeResourcePackageResponseBody {
	return s.Body
}

func (s *UpgradeResourcePackageResponse) SetHeaders(v map[string]*string) *UpgradeResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeResourcePackageResponse) SetStatusCode(v int32) *UpgradeResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeResourcePackageResponse) SetBody(v *UpgradeResourcePackageResponseBody) *UpgradeResourcePackageResponse {
	s.Body = v
	return s
}

func (s *UpgradeResourcePackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
