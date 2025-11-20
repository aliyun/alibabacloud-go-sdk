// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDriveSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDriveSpaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDriveSpaceResponseBody) *DeleteDriveSpaceResponse
	GetBody() *DeleteDriveSpaceResponseBody
}

type DeleteDriveSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDriveSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDriveSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDriveSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDriveSpaceResponse) GetBody() *DeleteDriveSpaceResponseBody {
	return s.Body
}

func (s *DeleteDriveSpaceResponse) SetHeaders(v map[string]*string) *DeleteDriveSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDriveSpaceResponse) SetStatusCode(v int32) *DeleteDriveSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDriveSpaceResponse) SetBody(v *DeleteDriveSpaceResponseBody) *DeleteDriveSpaceResponse {
	s.Body = v
	return s
}

func (s *DeleteDriveSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
