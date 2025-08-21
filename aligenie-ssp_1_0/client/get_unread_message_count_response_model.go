// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnreadMessageCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUnreadMessageCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUnreadMessageCountResponse
	GetStatusCode() *int32
	SetBody(v *GetUnreadMessageCountResponseBody) *GetUnreadMessageCountResponse
	GetBody() *GetUnreadMessageCountResponseBody
}

type GetUnreadMessageCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnreadMessageCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnreadMessageCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountResponse) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUnreadMessageCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnreadMessageCountResponse) GetBody() *GetUnreadMessageCountResponseBody {
	return s.Body
}

func (s *GetUnreadMessageCountResponse) SetHeaders(v map[string]*string) *GetUnreadMessageCountResponse {
	s.Headers = v
	return s
}

func (s *GetUnreadMessageCountResponse) SetStatusCode(v int32) *GetUnreadMessageCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnreadMessageCountResponse) SetBody(v *GetUnreadMessageCountResponseBody) *GetUnreadMessageCountResponse {
	s.Body = v
	return s
}

func (s *GetUnreadMessageCountResponse) Validate() error {
	return dara.Validate(s)
}
