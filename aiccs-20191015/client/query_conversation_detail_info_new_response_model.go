// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConversationDetailInfoNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConversationDetailInfoNewResponse
	GetStatusCode() *int32
	SetBody(v *QueryConversationDetailInfoNewResponseBody) *QueryConversationDetailInfoNewResponse
	GetBody() *QueryConversationDetailInfoNewResponseBody
}

type QueryConversationDetailInfoNewResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationDetailInfoNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationDetailInfoNewResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConversationDetailInfoNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConversationDetailInfoNewResponse) GetBody() *QueryConversationDetailInfoNewResponseBody {
	return s.Body
}

func (s *QueryConversationDetailInfoNewResponse) SetHeaders(v map[string]*string) *QueryConversationDetailInfoNewResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationDetailInfoNewResponse) SetStatusCode(v int32) *QueryConversationDetailInfoNewResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponse) SetBody(v *QueryConversationDetailInfoNewResponseBody) *QueryConversationDetailInfoNewResponse {
	s.Body = v
	return s
}

func (s *QueryConversationDetailInfoNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
