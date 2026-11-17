// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConversationShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableConversationShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableConversationShareResponse
	GetStatusCode() *int32
	SetBody(v *DisableConversationShareResponseBody) *DisableConversationShareResponse
	GetBody() *DisableConversationShareResponseBody
}

type DisableConversationShareResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableConversationShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableConversationShareResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableConversationShareResponse) GoString() string {
	return s.String()
}

func (s *DisableConversationShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableConversationShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableConversationShareResponse) GetBody() *DisableConversationShareResponseBody {
	return s.Body
}

func (s *DisableConversationShareResponse) SetHeaders(v map[string]*string) *DisableConversationShareResponse {
	s.Headers = v
	return s
}

func (s *DisableConversationShareResponse) SetStatusCode(v int32) *DisableConversationShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableConversationShareResponse) SetBody(v *DisableConversationShareResponseBody) *DisableConversationShareResponse {
	s.Body = v
	return s
}

func (s *DisableConversationShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
