// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePackageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePackageConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePackageConfigResponseBody) *DeleteLivePackageConfigResponse
	GetBody() *DeleteLivePackageConfigResponseBody
}

type DeleteLivePackageConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePackageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePackageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePackageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePackageConfigResponse) GetBody() *DeleteLivePackageConfigResponseBody {
	return s.Body
}

func (s *DeleteLivePackageConfigResponse) SetHeaders(v map[string]*string) *DeleteLivePackageConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePackageConfigResponse) SetStatusCode(v int32) *DeleteLivePackageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePackageConfigResponse) SetBody(v *DeleteLivePackageConfigResponseBody) *DeleteLivePackageConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePackageConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
