// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAndroidInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeAndroidInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeAndroidInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeAndroidInstanceGroupResponseBody) *UpgradeAndroidInstanceGroupResponse
	GetBody() *UpgradeAndroidInstanceGroupResponseBody
}

type UpgradeAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAndroidInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeAndroidInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeAndroidInstanceGroupResponse) GetBody() *UpgradeAndroidInstanceGroupResponseBody {
	return s.Body
}

func (s *UpgradeAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *UpgradeAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponse) SetStatusCode(v int32) *UpgradeAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponse) SetBody(v *UpgradeAndroidInstanceGroupResponseBody) *UpgradeAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
