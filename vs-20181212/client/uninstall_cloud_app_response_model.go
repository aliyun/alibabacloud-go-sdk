// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallCloudAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallCloudAppResponse
	GetStatusCode() *int32
	SetBody(v *UninstallCloudAppResponseBody) *UninstallCloudAppResponse
	GetBody() *UninstallCloudAppResponseBody
}

type UninstallCloudAppResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallCloudAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallCloudAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppResponse) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallCloudAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallCloudAppResponse) GetBody() *UninstallCloudAppResponseBody {
	return s.Body
}

func (s *UninstallCloudAppResponse) SetHeaders(v map[string]*string) *UninstallCloudAppResponse {
	s.Headers = v
	return s
}

func (s *UninstallCloudAppResponse) SetStatusCode(v int32) *UninstallCloudAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallCloudAppResponse) SetBody(v *UninstallCloudAppResponseBody) *UninstallCloudAppResponse {
	s.Body = v
	return s
}

func (s *UninstallCloudAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
