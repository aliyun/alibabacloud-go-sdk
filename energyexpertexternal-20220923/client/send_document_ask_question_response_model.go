// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDocumentAskQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendDocumentAskQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendDocumentAskQuestionResponse
	GetStatusCode() *int32
	SetBody(v *SendDocumentAskQuestionResponseBody) *SendDocumentAskQuestionResponse
	GetBody() *SendDocumentAskQuestionResponseBody
}

type SendDocumentAskQuestionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDocumentAskQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDocumentAskQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s SendDocumentAskQuestionResponse) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendDocumentAskQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendDocumentAskQuestionResponse) GetBody() *SendDocumentAskQuestionResponseBody {
	return s.Body
}

func (s *SendDocumentAskQuestionResponse) SetHeaders(v map[string]*string) *SendDocumentAskQuestionResponse {
	s.Headers = v
	return s
}

func (s *SendDocumentAskQuestionResponse) SetStatusCode(v int32) *SendDocumentAskQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDocumentAskQuestionResponse) SetBody(v *SendDocumentAskQuestionResponseBody) *SendDocumentAskQuestionResponse {
	s.Body = v
	return s
}

func (s *SendDocumentAskQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
