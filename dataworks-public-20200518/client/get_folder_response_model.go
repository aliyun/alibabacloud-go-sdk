// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFolderResponse
	GetStatusCode() *int32
	SetBody(v *GetFolderResponseBody) *GetFolderResponse
	GetBody() *GetFolderResponseBody
}

type GetFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponse) GoString() string {
	return s.String()
}

func (s *GetFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFolderResponse) GetBody() *GetFolderResponseBody {
	return s.Body
}

func (s *GetFolderResponse) SetHeaders(v map[string]*string) *GetFolderResponse {
	s.Headers = v
	return s
}

func (s *GetFolderResponse) SetStatusCode(v int32) *GetFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFolderResponse) SetBody(v *GetFolderResponseBody) *GetFolderResponse {
	s.Body = v
	return s
}

func (s *GetFolderResponse) Validate() error {
	return dara.Validate(s)
}
