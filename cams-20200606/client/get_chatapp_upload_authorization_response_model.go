// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappUploadAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappUploadAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappUploadAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappUploadAuthorizationResponseBody) *GetChatappUploadAuthorizationResponse
	GetBody() *GetChatappUploadAuthorizationResponseBody
}

type GetChatappUploadAuthorizationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappUploadAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappUploadAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappUploadAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappUploadAuthorizationResponse) GetBody() *GetChatappUploadAuthorizationResponseBody {
	return s.Body
}

func (s *GetChatappUploadAuthorizationResponse) SetHeaders(v map[string]*string) *GetChatappUploadAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetChatappUploadAuthorizationResponse) SetStatusCode(v int32) *GetChatappUploadAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponse) SetBody(v *GetChatappUploadAuthorizationResponseBody) *GetChatappUploadAuthorizationResponse {
	s.Body = v
	return s
}

func (s *GetChatappUploadAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
