// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAICoachTaskSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAICoachTaskSessionResponse
	GetStatusCode() *int32
	SetBody(v *ListAICoachTaskSessionResponseBody) *ListAICoachTaskSessionResponse
	GetBody() *ListAICoachTaskSessionResponseBody
}

type ListAICoachTaskSessionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICoachTaskSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAICoachTaskSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAICoachTaskSessionResponse) GetBody() *ListAICoachTaskSessionResponseBody {
	return s.Body
}

func (s *ListAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *ListAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *ListAICoachTaskSessionResponse) SetStatusCode(v int32) *ListAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICoachTaskSessionResponse) SetBody(v *ListAICoachTaskSessionResponseBody) *ListAICoachTaskSessionResponse {
	s.Body = v
	return s
}

func (s *ListAICoachTaskSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
