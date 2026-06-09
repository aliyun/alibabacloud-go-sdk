// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CutQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CutQuestionsResponse
	GetStatusCode() *int32
	SetBody(v *CutQuestionsResponseBody) *CutQuestionsResponse
	GetBody() *CutQuestionsResponseBody
}

type CutQuestionsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CutQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CutQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CutQuestionsResponse) GoString() string {
	return s.String()
}

func (s *CutQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CutQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CutQuestionsResponse) GetBody() *CutQuestionsResponseBody {
	return s.Body
}

func (s *CutQuestionsResponse) SetHeaders(v map[string]*string) *CutQuestionsResponse {
	s.Headers = v
	return s
}

func (s *CutQuestionsResponse) SetStatusCode(v int32) *CutQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CutQuestionsResponse) SetBody(v *CutQuestionsResponseBody) *CutQuestionsResponse {
	s.Body = v
	return s
}

func (s *CutQuestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
