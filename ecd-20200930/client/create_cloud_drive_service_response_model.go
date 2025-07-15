// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudDriveServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudDriveServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudDriveServiceResponseBody) *CreateCloudDriveServiceResponse
	GetBody() *CreateCloudDriveServiceResponseBody
}

type CreateCloudDriveServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudDriveServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudDriveServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudDriveServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudDriveServiceResponse) GetBody() *CreateCloudDriveServiceResponseBody {
	return s.Body
}

func (s *CreateCloudDriveServiceResponse) SetHeaders(v map[string]*string) *CreateCloudDriveServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudDriveServiceResponse) SetStatusCode(v int32) *CreateCloudDriveServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudDriveServiceResponse) SetBody(v *CreateCloudDriveServiceResponseBody) *CreateCloudDriveServiceResponse {
	s.Body = v
	return s
}

func (s *CreateCloudDriveServiceResponse) Validate() error {
	return dara.Validate(s)
}
