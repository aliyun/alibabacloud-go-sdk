// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeEnvironmentFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeEnvironmentFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeEnvironmentFeatureResponseBody) *UpgradeEnvironmentFeatureResponse
	GetBody() *UpgradeEnvironmentFeatureResponseBody
}

type UpgradeEnvironmentFeatureResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeEnvironmentFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeEnvironmentFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeEnvironmentFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeEnvironmentFeatureResponse) GetBody() *UpgradeEnvironmentFeatureResponseBody {
	return s.Body
}

func (s *UpgradeEnvironmentFeatureResponse) SetHeaders(v map[string]*string) *UpgradeEnvironmentFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpgradeEnvironmentFeatureResponse) SetStatusCode(v int32) *UpgradeEnvironmentFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeEnvironmentFeatureResponse) SetBody(v *UpgradeEnvironmentFeatureResponseBody) *UpgradeEnvironmentFeatureResponse {
	s.Body = v
	return s
}

func (s *UpgradeEnvironmentFeatureResponse) Validate() error {
	return dara.Validate(s)
}
