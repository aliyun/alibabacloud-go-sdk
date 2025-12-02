// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAnswerSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAnswerSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAnswerSampleResponse
	GetStatusCode() *int32
	SetBody(v *AddAnswerSampleResponseBody) *AddAnswerSampleResponse
	GetBody() *AddAnswerSampleResponseBody
}

type AddAnswerSampleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAnswerSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAnswerSampleResponse) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAnswerSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAnswerSampleResponse) GetBody() *AddAnswerSampleResponseBody {
	return s.Body
}

func (s *AddAnswerSampleResponse) SetHeaders(v map[string]*string) *AddAnswerSampleResponse {
	s.Headers = v
	return s
}

func (s *AddAnswerSampleResponse) SetStatusCode(v int32) *AddAnswerSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAnswerSampleResponse) SetBody(v *AddAnswerSampleResponseBody) *AddAnswerSampleResponse {
	s.Body = v
	return s
}

func (s *AddAnswerSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
