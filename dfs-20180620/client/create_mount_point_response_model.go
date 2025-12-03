// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMountPointResponse
	GetStatusCode() *int32
	SetBody(v *CreateMountPointResponseBody) *CreateMountPointResponse
	GetBody() *CreateMountPointResponseBody
}

type CreateMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMountPointResponse) GoString() string {
	return s.String()
}

func (s *CreateMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMountPointResponse) GetBody() *CreateMountPointResponseBody {
	return s.Body
}

func (s *CreateMountPointResponse) SetHeaders(v map[string]*string) *CreateMountPointResponse {
	s.Headers = v
	return s
}

func (s *CreateMountPointResponse) SetStatusCode(v int32) *CreateMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMountPointResponse) SetBody(v *CreateMountPointResponseBody) *CreateMountPointResponse {
	s.Body = v
	return s
}

func (s *CreateMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
