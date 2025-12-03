// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVscMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVscMountPointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVscMountPointResponseBody) *DeleteVscMountPointResponse
	GetBody() *DeleteVscMountPointResponseBody
}

type DeleteVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVscMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVscMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVscMountPointResponse) GetBody() *DeleteVscMountPointResponseBody {
	return s.Body
}

func (s *DeleteVscMountPointResponse) SetHeaders(v map[string]*string) *DeleteVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVscMountPointResponse) SetStatusCode(v int32) *DeleteVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVscMountPointResponse) SetBody(v *DeleteVscMountPointResponseBody) *DeleteVscMountPointResponse {
	s.Body = v
	return s
}

func (s *DeleteVscMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
