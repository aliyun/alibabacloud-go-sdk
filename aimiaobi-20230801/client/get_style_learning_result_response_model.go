// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleLearningResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStyleLearningResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStyleLearningResultResponse
	GetStatusCode() *int32
	SetBody(v *GetStyleLearningResultResponseBody) *GetStyleLearningResultResponse
	GetBody() *GetStyleLearningResultResponseBody
}

type GetStyleLearningResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStyleLearningResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStyleLearningResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultResponse) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStyleLearningResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStyleLearningResultResponse) GetBody() *GetStyleLearningResultResponseBody {
	return s.Body
}

func (s *GetStyleLearningResultResponse) SetHeaders(v map[string]*string) *GetStyleLearningResultResponse {
	s.Headers = v
	return s
}

func (s *GetStyleLearningResultResponse) SetStatusCode(v int32) *GetStyleLearningResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStyleLearningResultResponse) SetBody(v *GetStyleLearningResultResponseBody) *GetStyleLearningResultResponse {
	s.Body = v
	return s
}

func (s *GetStyleLearningResultResponse) Validate() error {
	return dara.Validate(s)
}
