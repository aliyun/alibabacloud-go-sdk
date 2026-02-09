// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunGenerateQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunGenerateQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunGenerateQuestionsResponse
	GetStatusCode() *int32
	SetId(v string) *RunGenerateQuestionsResponse
	GetId() *string
	SetEvent(v string) *RunGenerateQuestionsResponse
	GetEvent() *string
	SetBody(v *RunGenerateQuestionsResponseBody) *RunGenerateQuestionsResponse
	GetBody() *RunGenerateQuestionsResponseBody
}

type RunGenerateQuestionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunGenerateQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunGenerateQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponse) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunGenerateQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunGenerateQuestionsResponse) GetId() *string {
	return s.Id
}

func (s *RunGenerateQuestionsResponse) GetEvent() *string {
	return s.Event
}

func (s *RunGenerateQuestionsResponse) GetBody() *RunGenerateQuestionsResponseBody {
	return s.Body
}

func (s *RunGenerateQuestionsResponse) SetHeaders(v map[string]*string) *RunGenerateQuestionsResponse {
	s.Headers = v
	return s
}

func (s *RunGenerateQuestionsResponse) SetStatusCode(v int32) *RunGenerateQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *RunGenerateQuestionsResponse) SetId(v string) *RunGenerateQuestionsResponse {
	s.Id = &v
	return s
}

func (s *RunGenerateQuestionsResponse) SetEvent(v string) *RunGenerateQuestionsResponse {
	s.Event = &v
	return s
}

func (s *RunGenerateQuestionsResponse) SetBody(v *RunGenerateQuestionsResponseBody) *RunGenerateQuestionsResponse {
	s.Body = v
	return s
}

func (s *RunGenerateQuestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
