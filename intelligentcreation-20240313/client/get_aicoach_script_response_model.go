// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachScriptResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachScriptResponseBody) *GetAICoachScriptResponse
	GetBody() *GetAICoachScriptResponseBody
}

type GetAICoachScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachScriptResponse) GetBody() *GetAICoachScriptResponseBody {
	return s.Body
}

func (s *GetAICoachScriptResponse) SetHeaders(v map[string]*string) *GetAICoachScriptResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachScriptResponse) SetStatusCode(v int32) *GetAICoachScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachScriptResponse) SetBody(v *GetAICoachScriptResponseBody) *GetAICoachScriptResponse {
	s.Body = v
	return s
}

func (s *GetAICoachScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
