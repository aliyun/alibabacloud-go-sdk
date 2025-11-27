// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverMajorVersionUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchOverMajorVersionUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchOverMajorVersionUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *SwitchOverMajorVersionUpgradeResponseBody) *SwitchOverMajorVersionUpgradeResponse
	GetBody() *SwitchOverMajorVersionUpgradeResponseBody
}

type SwitchOverMajorVersionUpgradeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchOverMajorVersionUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchOverMajorVersionUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverMajorVersionUpgradeResponse) GoString() string {
	return s.String()
}

func (s *SwitchOverMajorVersionUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchOverMajorVersionUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchOverMajorVersionUpgradeResponse) GetBody() *SwitchOverMajorVersionUpgradeResponseBody {
	return s.Body
}

func (s *SwitchOverMajorVersionUpgradeResponse) SetHeaders(v map[string]*string) *SwitchOverMajorVersionUpgradeResponse {
	s.Headers = v
	return s
}

func (s *SwitchOverMajorVersionUpgradeResponse) SetStatusCode(v int32) *SwitchOverMajorVersionUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeResponse) SetBody(v *SwitchOverMajorVersionUpgradeResponseBody) *SwitchOverMajorVersionUpgradeResponse {
	s.Body = v
	return s
}

func (s *SwitchOverMajorVersionUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
