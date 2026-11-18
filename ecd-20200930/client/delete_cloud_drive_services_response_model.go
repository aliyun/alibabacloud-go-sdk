// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudDriveServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudDriveServicesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudDriveServicesResponseBody) *DeleteCloudDriveServicesResponse
	GetBody() *DeleteCloudDriveServicesResponseBody
}

type DeleteCloudDriveServicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudDriveServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudDriveServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveServicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudDriveServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudDriveServicesResponse) GetBody() *DeleteCloudDriveServicesResponseBody {
	return s.Body
}

func (s *DeleteCloudDriveServicesResponse) SetHeaders(v map[string]*string) *DeleteCloudDriveServicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudDriveServicesResponse) SetStatusCode(v int32) *DeleteCloudDriveServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudDriveServicesResponse) SetBody(v *DeleteCloudDriveServicesResponseBody) *DeleteCloudDriveServicesResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudDriveServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
