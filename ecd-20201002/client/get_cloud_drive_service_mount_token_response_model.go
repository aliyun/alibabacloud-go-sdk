// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudDriveServiceMountTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudDriveServiceMountTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudDriveServiceMountTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudDriveServiceMountTokenResponseBody) *GetCloudDriveServiceMountTokenResponse
	GetBody() *GetCloudDriveServiceMountTokenResponseBody
}

type GetCloudDriveServiceMountTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudDriveServiceMountTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudDriveServiceMountTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudDriveServiceMountTokenResponse) GetBody() *GetCloudDriveServiceMountTokenResponseBody {
	return s.Body
}

func (s *GetCloudDriveServiceMountTokenResponse) SetHeaders(v map[string]*string) *GetCloudDriveServiceMountTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetStatusCode(v int32) *GetCloudDriveServiceMountTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetBody(v *GetCloudDriveServiceMountTokenResponseBody) *GetCloudDriveServiceMountTokenResponse {
	s.Body = v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) Validate() error {
	return dara.Validate(s)
}
