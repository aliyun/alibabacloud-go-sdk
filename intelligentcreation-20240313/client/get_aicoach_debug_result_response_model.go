// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachDebugResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachDebugResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachDebugResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachDebugResultResponseBody) *GetAICoachDebugResultResponse
	GetBody() *GetAICoachDebugResultResponseBody
}

type GetAICoachDebugResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachDebugResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachDebugResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachDebugResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachDebugResultResponse) GetBody() *GetAICoachDebugResultResponseBody {
	return s.Body
}

func (s *GetAICoachDebugResultResponse) SetHeaders(v map[string]*string) *GetAICoachDebugResultResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachDebugResultResponse) SetStatusCode(v int32) *GetAICoachDebugResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachDebugResultResponse) SetBody(v *GetAICoachDebugResultResponseBody) *GetAICoachDebugResultResponse {
	s.Body = v
	return s
}

func (s *GetAICoachDebugResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
