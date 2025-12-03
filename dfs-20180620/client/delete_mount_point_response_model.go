// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMountPointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMountPointResponseBody) *DeleteMountPointResponse
	GetBody() *DeleteMountPointResponseBody
}

type DeleteMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMountPointResponse) GetBody() *DeleteMountPointResponseBody {
	return s.Body
}

func (s *DeleteMountPointResponse) SetHeaders(v map[string]*string) *DeleteMountPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteMountPointResponse) SetStatusCode(v int32) *DeleteMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMountPointResponse) SetBody(v *DeleteMountPointResponseBody) *DeleteMountPointResponse {
	s.Body = v
	return s
}

func (s *DeleteMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
