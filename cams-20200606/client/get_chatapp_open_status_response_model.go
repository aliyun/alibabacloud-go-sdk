// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappOpenStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappOpenStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappOpenStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappOpenStatusResponseBody) *GetChatappOpenStatusResponse
	GetBody() *GetChatappOpenStatusResponseBody
}

type GetChatappOpenStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappOpenStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetChatappOpenStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappOpenStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappOpenStatusResponse) GetBody() *GetChatappOpenStatusResponseBody {
	return s.Body
}

func (s *GetChatappOpenStatusResponse) SetHeaders(v map[string]*string) *GetChatappOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetChatappOpenStatusResponse) SetStatusCode(v int32) *GetChatappOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappOpenStatusResponse) SetBody(v *GetChatappOpenStatusResponseBody) *GetChatappOpenStatusResponse {
	s.Body = v
	return s
}

func (s *GetChatappOpenStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
