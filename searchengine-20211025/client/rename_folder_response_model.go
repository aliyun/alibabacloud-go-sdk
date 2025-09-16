// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameFolderResponse
	GetStatusCode() *int32
	SetBody(v *RenameFolderResponseBody) *RenameFolderResponse
	GetBody() *RenameFolderResponseBody
}

type RenameFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameFolderResponse) GoString() string {
	return s.String()
}

func (s *RenameFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameFolderResponse) GetBody() *RenameFolderResponseBody {
	return s.Body
}

func (s *RenameFolderResponse) SetHeaders(v map[string]*string) *RenameFolderResponse {
	s.Headers = v
	return s
}

func (s *RenameFolderResponse) SetStatusCode(v int32) *RenameFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameFolderResponse) SetBody(v *RenameFolderResponseBody) *RenameFolderResponse {
	s.Body = v
	return s
}

func (s *RenameFolderResponse) Validate() error {
	return dara.Validate(s)
}
