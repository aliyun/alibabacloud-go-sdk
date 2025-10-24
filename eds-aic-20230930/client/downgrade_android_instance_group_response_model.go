// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeAndroidInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradeAndroidInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradeAndroidInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DowngradeAndroidInstanceGroupResponseBody) *DowngradeAndroidInstanceGroupResponse
	GetBody() *DowngradeAndroidInstanceGroupResponseBody
}

type DowngradeAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradeAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradeAndroidInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradeAndroidInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradeAndroidInstanceGroupResponse) GetBody() *DowngradeAndroidInstanceGroupResponseBody {
	return s.Body
}

func (s *DowngradeAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *DowngradeAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponse) SetStatusCode(v int32) *DowngradeAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponse) SetBody(v *DowngradeAndroidInstanceGroupResponseBody) *DowngradeAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
