// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFolderResponse
	GetStatusCode() *int32
	SetBody(v *CreateFolderResponseBody) *CreateFolderResponse
	GetBody() *CreateFolderResponseBody
}

type CreateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFolderResponse) GetBody() *CreateFolderResponseBody {
	return s.Body
}

func (s *CreateFolderResponse) SetHeaders(v map[string]*string) *CreateFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateFolderResponse) SetStatusCode(v int32) *CreateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFolderResponse) SetBody(v *CreateFolderResponseBody) *CreateFolderResponse {
	s.Body = v
	return s
}

func (s *CreateFolderResponse) Validate() error {
	return dara.Validate(s)
}
