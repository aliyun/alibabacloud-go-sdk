// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *GetFileSystemResponseBody) *GetFileSystemResponse
	GetBody() *GetFileSystemResponseBody
}

type GetFileSystemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileSystemResponse) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileSystemResponse) GetBody() *GetFileSystemResponseBody {
	return s.Body
}

func (s *GetFileSystemResponse) SetHeaders(v map[string]*string) *GetFileSystemResponse {
	s.Headers = v
	return s
}

func (s *GetFileSystemResponse) SetStatusCode(v int32) *GetFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileSystemResponse) SetBody(v *GetFileSystemResponseBody) *GetFileSystemResponse {
	s.Body = v
	return s
}

func (s *GetFileSystemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
