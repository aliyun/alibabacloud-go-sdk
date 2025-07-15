// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConversationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConversationDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetConversationDetailResponseBody) *GetConversationDetailResponse
	GetBody() *GetConversationDetailResponseBody
}

type GetConversationDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConversationDetailResponse) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConversationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConversationDetailResponse) GetBody() *GetConversationDetailResponseBody {
	return s.Body
}

func (s *GetConversationDetailResponse) SetHeaders(v map[string]*string) *GetConversationDetailResponse {
	s.Headers = v
	return s
}

func (s *GetConversationDetailResponse) SetStatusCode(v int32) *GetConversationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationDetailResponse) SetBody(v *GetConversationDetailResponseBody) *GetConversationDetailResponse {
	s.Body = v
	return s
}

func (s *GetConversationDetailResponse) Validate() error {
	return dara.Validate(s)
}
