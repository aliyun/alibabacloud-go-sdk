// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFolderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFolderResponseBody) *UpdateFolderResponse
	GetBody() *UpdateFolderResponseBody
}

type UpdateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFolderResponse) GetBody() *UpdateFolderResponseBody {
	return s.Body
}

func (s *UpdateFolderResponse) SetHeaders(v map[string]*string) *UpdateFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateFolderResponse) SetStatusCode(v int32) *UpdateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFolderResponse) SetBody(v *UpdateFolderResponseBody) *UpdateFolderResponse {
	s.Body = v
	return s
}

func (s *UpdateFolderResponse) Validate() error {
	return dara.Validate(s)
}
