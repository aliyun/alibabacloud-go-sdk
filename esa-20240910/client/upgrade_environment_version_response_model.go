// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeEnvironmentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeEnvironmentVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeEnvironmentVersionResponseBody) *UpgradeEnvironmentVersionResponse
	GetBody() *UpgradeEnvironmentVersionResponseBody
}

type UpgradeEnvironmentVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeEnvironmentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeEnvironmentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeEnvironmentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeEnvironmentVersionResponse) GetBody() *UpgradeEnvironmentVersionResponseBody {
	return s.Body
}

func (s *UpgradeEnvironmentVersionResponse) SetHeaders(v map[string]*string) *UpgradeEnvironmentVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeEnvironmentVersionResponse) SetStatusCode(v int32) *UpgradeEnvironmentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeEnvironmentVersionResponse) SetBody(v *UpgradeEnvironmentVersionResponseBody) *UpgradeEnvironmentVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeEnvironmentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
