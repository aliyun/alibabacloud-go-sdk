// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFolderResponse
	GetStatusCode() *int32
	SetBody(v *AddFolderResponseBody) *AddFolderResponse
	GetBody() *AddFolderResponseBody
}

type AddFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFolderResponse) GoString() string {
	return s.String()
}

func (s *AddFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFolderResponse) GetBody() *AddFolderResponseBody {
	return s.Body
}

func (s *AddFolderResponse) SetHeaders(v map[string]*string) *AddFolderResponse {
	s.Headers = v
	return s
}

func (s *AddFolderResponse) SetStatusCode(v int32) *AddFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFolderResponse) SetBody(v *AddFolderResponseBody) *AddFolderResponse {
	s.Body = v
	return s
}

func (s *AddFolderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
