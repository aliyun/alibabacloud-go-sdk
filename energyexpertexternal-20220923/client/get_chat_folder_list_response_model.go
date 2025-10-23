// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFolderListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatFolderListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatFolderListResponse
	GetStatusCode() *int32
	SetBody(v *GetChatFolderListResponseBody) *GetChatFolderListResponse
	GetBody() *GetChatFolderListResponseBody
}

type GetChatFolderListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatFolderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatFolderListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatFolderListResponse) GoString() string {
	return s.String()
}

func (s *GetChatFolderListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatFolderListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatFolderListResponse) GetBody() *GetChatFolderListResponseBody {
	return s.Body
}

func (s *GetChatFolderListResponse) SetHeaders(v map[string]*string) *GetChatFolderListResponse {
	s.Headers = v
	return s
}

func (s *GetChatFolderListResponse) SetStatusCode(v int32) *GetChatFolderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatFolderListResponse) SetBody(v *GetChatFolderListResponseBody) *GetChatFolderListResponse {
	s.Body = v
	return s
}

func (s *GetChatFolderListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
