// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToggleAutoMinorVersionUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ToggleAutoMinorVersionUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ToggleAutoMinorVersionUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *ToggleAutoMinorVersionUpgradeResponseBody) *ToggleAutoMinorVersionUpgradeResponse
	GetBody() *ToggleAutoMinorVersionUpgradeResponseBody
}

type ToggleAutoMinorVersionUpgradeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToggleAutoMinorVersionUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ToggleAutoMinorVersionUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s ToggleAutoMinorVersionUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ToggleAutoMinorVersionUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ToggleAutoMinorVersionUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ToggleAutoMinorVersionUpgradeResponse) GetBody() *ToggleAutoMinorVersionUpgradeResponseBody {
	return s.Body
}

func (s *ToggleAutoMinorVersionUpgradeResponse) SetHeaders(v map[string]*string) *ToggleAutoMinorVersionUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponse) SetStatusCode(v int32) *ToggleAutoMinorVersionUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponse) SetBody(v *ToggleAutoMinorVersionUpgradeResponseBody) *ToggleAutoMinorVersionUpgradeResponse {
	s.Body = v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
