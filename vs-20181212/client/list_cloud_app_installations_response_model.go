// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppInstallationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAppInstallationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAppInstallationsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAppInstallationsResponseBody) *ListCloudAppInstallationsResponse
	GetBody() *ListCloudAppInstallationsResponseBody
}

type ListCloudAppInstallationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAppInstallationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAppInstallationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppInstallationsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAppInstallationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAppInstallationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAppInstallationsResponse) GetBody() *ListCloudAppInstallationsResponseBody {
	return s.Body
}

func (s *ListCloudAppInstallationsResponse) SetHeaders(v map[string]*string) *ListCloudAppInstallationsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAppInstallationsResponse) SetStatusCode(v int32) *ListCloudAppInstallationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAppInstallationsResponse) SetBody(v *ListCloudAppInstallationsResponseBody) *ListCloudAppInstallationsResponse {
	s.Body = v
	return s
}

func (s *ListCloudAppInstallationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
