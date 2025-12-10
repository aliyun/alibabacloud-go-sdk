// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFolderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFolderResponseBody) *DeleteFolderResponse
	GetBody() *DeleteFolderResponseBody
}

type DeleteFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFolderResponse) GetBody() *DeleteFolderResponseBody {
	return s.Body
}

func (s *DeleteFolderResponse) SetHeaders(v map[string]*string) *DeleteFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteFolderResponse) SetStatusCode(v int32) *DeleteFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFolderResponse) SetBody(v *DeleteFolderResponseBody) *DeleteFolderResponse {
	s.Body = v
	return s
}

func (s *DeleteFolderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
