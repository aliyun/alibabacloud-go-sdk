// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGenerateAICoachScriptTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGenerateAICoachScriptTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGenerateAICoachScriptTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateGenerateAICoachScriptTaskResponseBody) *CreateGenerateAICoachScriptTaskResponse
	GetBody() *CreateGenerateAICoachScriptTaskResponseBody
}

type CreateGenerateAICoachScriptTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGenerateAICoachScriptTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGenerateAICoachScriptTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGenerateAICoachScriptTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGenerateAICoachScriptTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGenerateAICoachScriptTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGenerateAICoachScriptTaskResponse) GetBody() *CreateGenerateAICoachScriptTaskResponseBody {
	return s.Body
}

func (s *CreateGenerateAICoachScriptTaskResponse) SetHeaders(v map[string]*string) *CreateGenerateAICoachScriptTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponse) SetStatusCode(v int32) *CreateGenerateAICoachScriptTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponse) SetBody(v *CreateGenerateAICoachScriptTaskResponseBody) *CreateGenerateAICoachScriptTaskResponse {
	s.Body = v
	return s
}

func (s *CreateGenerateAICoachScriptTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
