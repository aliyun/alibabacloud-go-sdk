// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVscMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVscMountPointResponse
	GetStatusCode() *int32
	SetBody(v *CreateVscMountPointResponseBody) *CreateVscMountPointResponse
	GetBody() *CreateVscMountPointResponseBody
}

type CreateVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVscMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVscMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVscMountPointResponse) GetBody() *CreateVscMountPointResponseBody {
	return s.Body
}

func (s *CreateVscMountPointResponse) SetHeaders(v map[string]*string) *CreateVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *CreateVscMountPointResponse) SetStatusCode(v int32) *CreateVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVscMountPointResponse) SetBody(v *CreateVscMountPointResponseBody) *CreateVscMountPointResponse {
	s.Body = v
	return s
}

func (s *CreateVscMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
