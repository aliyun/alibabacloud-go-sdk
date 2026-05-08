// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishAICoachTaskSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishAICoachTaskSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishAICoachTaskSessionResponse
	GetStatusCode() *int32
	SetBody(v *FinishAICoachTaskSessionResponseBody) *FinishAICoachTaskSessionResponse
	GetBody() *FinishAICoachTaskSessionResponseBody
}

type FinishAICoachTaskSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishAICoachTaskSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishAICoachTaskSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishAICoachTaskSessionResponse) GetBody() *FinishAICoachTaskSessionResponseBody {
	return s.Body
}

func (s *FinishAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *FinishAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *FinishAICoachTaskSessionResponse) SetStatusCode(v int32) *FinishAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishAICoachTaskSessionResponse) SetBody(v *FinishAICoachTaskSessionResponseBody) *FinishAICoachTaskSessionResponse {
	s.Body = v
	return s
}

func (s *FinishAICoachTaskSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
