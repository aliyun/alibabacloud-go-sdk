// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMountPointResponse
	GetStatusCode() *int32
	SetBody(v *GetMountPointResponseBody) *GetMountPointResponse
	GetBody() *GetMountPointResponseBody
}

type GetMountPointResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMountPointResponse) GoString() string {
	return s.String()
}

func (s *GetMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMountPointResponse) GetBody() *GetMountPointResponseBody {
	return s.Body
}

func (s *GetMountPointResponse) SetHeaders(v map[string]*string) *GetMountPointResponse {
	s.Headers = v
	return s
}

func (s *GetMountPointResponse) SetStatusCode(v int32) *GetMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMountPointResponse) SetBody(v *GetMountPointResponseBody) *GetMountPointResponse {
	s.Body = v
	return s
}

func (s *GetMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
