// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEvaluationTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEvaluationTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *GetEvaluationTemplatesResponseBody) *GetEvaluationTemplatesResponse
	GetBody() *GetEvaluationTemplatesResponseBody
}

type GetEvaluationTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluationTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluationTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEvaluationTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEvaluationTemplatesResponse) GetBody() *GetEvaluationTemplatesResponseBody {
	return s.Body
}

func (s *GetEvaluationTemplatesResponse) SetHeaders(v map[string]*string) *GetEvaluationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluationTemplatesResponse) SetStatusCode(v int32) *GetEvaluationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluationTemplatesResponse) SetBody(v *GetEvaluationTemplatesResponseBody) *GetEvaluationTemplatesResponse {
	s.Body = v
	return s
}

func (s *GetEvaluationTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
