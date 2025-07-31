// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnswerSampleByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAnswerSampleByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAnswerSampleByPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryAnswerSampleByPageResponseBody) *QueryAnswerSampleByPageResponse
	GetBody() *QueryAnswerSampleByPageResponseBody
}

type QueryAnswerSampleByPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAnswerSampleByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAnswerSampleByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAnswerSampleByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAnswerSampleByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAnswerSampleByPageResponse) GetBody() *QueryAnswerSampleByPageResponseBody {
	return s.Body
}

func (s *QueryAnswerSampleByPageResponse) SetHeaders(v map[string]*string) *QueryAnswerSampleByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAnswerSampleByPageResponse) SetStatusCode(v int32) *QueryAnswerSampleByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAnswerSampleByPageResponse) SetBody(v *QueryAnswerSampleByPageResponseBody) *QueryAnswerSampleByPageResponse {
	s.Body = v
	return s
}

func (s *QueryAnswerSampleByPageResponse) Validate() error {
	return dara.Validate(s)
}
