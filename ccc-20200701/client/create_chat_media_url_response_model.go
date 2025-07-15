// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatMediaUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatMediaUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatMediaUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatMediaUrlResponseBody) *CreateChatMediaUrlResponse
	GetBody() *CreateChatMediaUrlResponseBody
}

type CreateChatMediaUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatMediaUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatMediaUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatMediaUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateChatMediaUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatMediaUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatMediaUrlResponse) GetBody() *CreateChatMediaUrlResponseBody {
	return s.Body
}

func (s *CreateChatMediaUrlResponse) SetHeaders(v map[string]*string) *CreateChatMediaUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateChatMediaUrlResponse) SetStatusCode(v int32) *CreateChatMediaUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatMediaUrlResponse) SetBody(v *CreateChatMediaUrlResponseBody) *CreateChatMediaUrlResponse {
	s.Body = v
	return s
}

func (s *CreateChatMediaUrlResponse) Validate() error {
	return dara.Validate(s)
}
