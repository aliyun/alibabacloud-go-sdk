// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConversationDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConversationDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryConversationDetailInfoResponseBody) *QueryConversationDetailInfoResponse
	GetBody() *QueryConversationDetailInfoResponseBody
}

type QueryConversationDetailInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConversationDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConversationDetailInfoResponse) GetBody() *QueryConversationDetailInfoResponseBody {
	return s.Body
}

func (s *QueryConversationDetailInfoResponse) SetHeaders(v map[string]*string) *QueryConversationDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationDetailInfoResponse) SetStatusCode(v int32) *QueryConversationDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationDetailInfoResponse) SetBody(v *QueryConversationDetailInfoResponseBody) *QueryConversationDetailInfoResponse {
	s.Body = v
	return s
}

func (s *QueryConversationDetailInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
