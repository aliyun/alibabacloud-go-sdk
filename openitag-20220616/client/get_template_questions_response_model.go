// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateQuestionsResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateQuestionsResponseBody) *GetTemplateQuestionsResponse
	GetBody() *GetTemplateQuestionsResponseBody
}

type GetTemplateQuestionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateQuestionsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateQuestionsResponse) GetBody() *GetTemplateQuestionsResponseBody {
	return s.Body
}

func (s *GetTemplateQuestionsResponse) SetHeaders(v map[string]*string) *GetTemplateQuestionsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateQuestionsResponse) SetStatusCode(v int32) *GetTemplateQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateQuestionsResponse) SetBody(v *GetTemplateQuestionsResponseBody) *GetTemplateQuestionsResponse {
	s.Body = v
	return s
}

func (s *GetTemplateQuestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
