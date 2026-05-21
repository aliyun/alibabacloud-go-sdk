// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeReleaseVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUpgradeReleaseVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUpgradeReleaseVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUpgradeReleaseVersionsResponseBody) *ListUpgradeReleaseVersionsResponse
	GetBody() *ListUpgradeReleaseVersionsResponseBody
}

type ListUpgradeReleaseVersionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUpgradeReleaseVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUpgradeReleaseVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeReleaseVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListUpgradeReleaseVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUpgradeReleaseVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUpgradeReleaseVersionsResponse) GetBody() *ListUpgradeReleaseVersionsResponseBody {
	return s.Body
}

func (s *ListUpgradeReleaseVersionsResponse) SetHeaders(v map[string]*string) *ListUpgradeReleaseVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListUpgradeReleaseVersionsResponse) SetStatusCode(v int32) *ListUpgradeReleaseVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponse) SetBody(v *ListUpgradeReleaseVersionsResponseBody) *ListUpgradeReleaseVersionsResponse {
	s.Body = v
	return s
}

func (s *ListUpgradeReleaseVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
