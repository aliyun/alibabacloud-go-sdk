// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachScriptPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAICoachScriptPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAICoachScriptPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAICoachScriptPageResponseBody) *ListAICoachScriptPageResponse
	GetBody() *ListAICoachScriptPageResponseBody
}

type ListAICoachScriptPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICoachScriptPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICoachScriptPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachScriptPageResponse) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAICoachScriptPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAICoachScriptPageResponse) GetBody() *ListAICoachScriptPageResponseBody {
	return s.Body
}

func (s *ListAICoachScriptPageResponse) SetHeaders(v map[string]*string) *ListAICoachScriptPageResponse {
	s.Headers = v
	return s
}

func (s *ListAICoachScriptPageResponse) SetStatusCode(v int32) *ListAICoachScriptPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICoachScriptPageResponse) SetBody(v *ListAICoachScriptPageResponseBody) *ListAICoachScriptPageResponse {
	s.Body = v
	return s
}

func (s *ListAICoachScriptPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
