// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptGenerateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachScriptGenerateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachScriptGenerateTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachScriptGenerateTaskResponseBody) *GetAICoachScriptGenerateTaskResponse
	GetBody() *GetAICoachScriptGenerateTaskResponseBody
}

type GetAICoachScriptGenerateTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachScriptGenerateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachScriptGenerateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptGenerateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptGenerateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachScriptGenerateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachScriptGenerateTaskResponse) GetBody() *GetAICoachScriptGenerateTaskResponseBody {
	return s.Body
}

func (s *GetAICoachScriptGenerateTaskResponse) SetHeaders(v map[string]*string) *GetAICoachScriptGenerateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponse) SetStatusCode(v int32) *GetAICoachScriptGenerateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponse) SetBody(v *GetAICoachScriptGenerateTaskResponseBody) *GetAICoachScriptGenerateTaskResponse {
	s.Body = v
	return s
}

func (s *GetAICoachScriptGenerateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
