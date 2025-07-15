// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindNasFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndBindNasFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndBindNasFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndBindNasFileSystemResponseBody) *CreateAndBindNasFileSystemResponse
	GetBody() *CreateAndBindNasFileSystemResponseBody
}

type CreateAndBindNasFileSystemResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndBindNasFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndBindNasFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindNasFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateAndBindNasFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndBindNasFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndBindNasFileSystemResponse) GetBody() *CreateAndBindNasFileSystemResponseBody {
	return s.Body
}

func (s *CreateAndBindNasFileSystemResponse) SetHeaders(v map[string]*string) *CreateAndBindNasFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateAndBindNasFileSystemResponse) SetStatusCode(v int32) *CreateAndBindNasFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndBindNasFileSystemResponse) SetBody(v *CreateAndBindNasFileSystemResponseBody) *CreateAndBindNasFileSystemResponse {
	s.Body = v
	return s
}

func (s *CreateAndBindNasFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
