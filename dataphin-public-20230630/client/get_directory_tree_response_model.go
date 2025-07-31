// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectoryTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectoryTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectoryTreeResponseBody) *GetDirectoryTreeResponse
	GetBody() *GetDirectoryTreeResponseBody
}

type GetDirectoryTreeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectoryTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectoryTreeResponse) GetBody() *GetDirectoryTreeResponseBody {
	return s.Body
}

func (s *GetDirectoryTreeResponse) SetHeaders(v map[string]*string) *GetDirectoryTreeResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryTreeResponse) SetStatusCode(v int32) *GetDirectoryTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryTreeResponse) SetBody(v *GetDirectoryTreeResponseBody) *GetDirectoryTreeResponse {
	s.Body = v
	return s
}

func (s *GetDirectoryTreeResponse) Validate() error {
	return dara.Validate(s)
}
