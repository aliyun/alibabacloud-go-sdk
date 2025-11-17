// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCubeDataLevelPermissionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCubeDataLevelPermissionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCubeDataLevelPermissionConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListCubeDataLevelPermissionConfigResponseBody) *ListCubeDataLevelPermissionConfigResponse
	GetBody() *ListCubeDataLevelPermissionConfigResponseBody
}

type ListCubeDataLevelPermissionConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCubeDataLevelPermissionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCubeDataLevelPermissionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCubeDataLevelPermissionConfigResponse) GoString() string {
	return s.String()
}

func (s *ListCubeDataLevelPermissionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCubeDataLevelPermissionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCubeDataLevelPermissionConfigResponse) GetBody() *ListCubeDataLevelPermissionConfigResponseBody {
	return s.Body
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetHeaders(v map[string]*string) *ListCubeDataLevelPermissionConfigResponse {
	s.Headers = v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetStatusCode(v int32) *ListCubeDataLevelPermissionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetBody(v *ListCubeDataLevelPermissionConfigResponseBody) *ListCubeDataLevelPermissionConfigResponse {
	s.Body = v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
