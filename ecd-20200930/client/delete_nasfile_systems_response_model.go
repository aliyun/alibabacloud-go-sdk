// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNASFileSystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNASFileSystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNASFileSystemsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNASFileSystemsResponseBody) *DeleteNASFileSystemsResponse
	GetBody() *DeleteNASFileSystemsResponseBody
}

type DeleteNASFileSystemsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNASFileSystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNASFileSystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNASFileSystemsResponse) GetBody() *DeleteNASFileSystemsResponseBody {
	return s.Body
}

func (s *DeleteNASFileSystemsResponse) SetHeaders(v map[string]*string) *DeleteNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNASFileSystemsResponse) SetStatusCode(v int32) *DeleteNASFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNASFileSystemsResponse) SetBody(v *DeleteNASFileSystemsResponseBody) *DeleteNASFileSystemsResponse {
	s.Body = v
	return s
}

func (s *DeleteNASFileSystemsResponse) Validate() error {
	return dara.Validate(s)
}
