// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNASFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNASFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNASFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *CreateNASFileSystemResponseBody) *CreateNASFileSystemResponse
	GetBody() *CreateNASFileSystemResponseBody
}

type CreateNASFileSystemResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNASFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNASFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNASFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNASFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNASFileSystemResponse) GetBody() *CreateNASFileSystemResponseBody {
	return s.Body
}

func (s *CreateNASFileSystemResponse) SetHeaders(v map[string]*string) *CreateNASFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateNASFileSystemResponse) SetStatusCode(v int32) *CreateNASFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNASFileSystemResponse) SetBody(v *CreateNASFileSystemResponseBody) *CreateNASFileSystemResponse {
	s.Body = v
	return s
}

func (s *CreateNASFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
