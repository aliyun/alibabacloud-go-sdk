// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAICoachTaskPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAICoachTaskPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAICoachTaskPageResponseBody) *ListAICoachTaskPageResponse
	GetBody() *ListAICoachTaskPageResponseBody
}

type ListAICoachTaskPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICoachTaskPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICoachTaskPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskPageResponse) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAICoachTaskPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAICoachTaskPageResponse) GetBody() *ListAICoachTaskPageResponseBody {
	return s.Body
}

func (s *ListAICoachTaskPageResponse) SetHeaders(v map[string]*string) *ListAICoachTaskPageResponse {
	s.Headers = v
	return s
}

func (s *ListAICoachTaskPageResponse) SetStatusCode(v int32) *ListAICoachTaskPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICoachTaskPageResponse) SetBody(v *ListAICoachTaskPageResponseBody) *ListAICoachTaskPageResponse {
	s.Body = v
	return s
}

func (s *ListAICoachTaskPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
