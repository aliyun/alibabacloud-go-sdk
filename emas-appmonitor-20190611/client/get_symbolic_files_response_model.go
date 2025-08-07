// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSymbolicFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSymbolicFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSymbolicFilesResponse
	GetStatusCode() *int32
	SetBody(v *GetSymbolicFilesResponseBody) *GetSymbolicFilesResponse
	GetBody() *GetSymbolicFilesResponseBody
}

type GetSymbolicFilesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSymbolicFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSymbolicFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSymbolicFilesResponse) GoString() string {
	return s.String()
}

func (s *GetSymbolicFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSymbolicFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSymbolicFilesResponse) GetBody() *GetSymbolicFilesResponseBody {
	return s.Body
}

func (s *GetSymbolicFilesResponse) SetHeaders(v map[string]*string) *GetSymbolicFilesResponse {
	s.Headers = v
	return s
}

func (s *GetSymbolicFilesResponse) SetStatusCode(v int32) *GetSymbolicFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSymbolicFilesResponse) SetBody(v *GetSymbolicFilesResponseBody) *GetSymbolicFilesResponse {
	s.Body = v
	return s
}

func (s *GetSymbolicFilesResponse) Validate() error {
	return dara.Validate(s)
}
