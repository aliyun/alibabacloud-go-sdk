// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAICoachScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineAICoachScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineAICoachScriptResponse
	GetStatusCode() *int32
	SetBody(v *OfflineAICoachScriptResponseBody) *OfflineAICoachScriptResponse
	GetBody() *OfflineAICoachScriptResponseBody
}

type OfflineAICoachScriptResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineAICoachScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineAICoachScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineAICoachScriptResponse) GoString() string {
	return s.String()
}

func (s *OfflineAICoachScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineAICoachScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineAICoachScriptResponse) GetBody() *OfflineAICoachScriptResponseBody {
	return s.Body
}

func (s *OfflineAICoachScriptResponse) SetHeaders(v map[string]*string) *OfflineAICoachScriptResponse {
	s.Headers = v
	return s
}

func (s *OfflineAICoachScriptResponse) SetStatusCode(v int32) *OfflineAICoachScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineAICoachScriptResponse) SetBody(v *OfflineAICoachScriptResponseBody) *OfflineAICoachScriptResponse {
	s.Body = v
	return s
}

func (s *OfflineAICoachScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
