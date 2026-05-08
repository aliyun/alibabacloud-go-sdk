// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAICoachTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAICoachTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAICoachTaskResponseBody) *CreateAICoachTaskResponse
	GetBody() *CreateAICoachTaskResponseBody
}

type CreateAICoachTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICoachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICoachTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAICoachTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAICoachTaskResponse) GetBody() *CreateAICoachTaskResponseBody {
	return s.Body
}

func (s *CreateAICoachTaskResponse) SetHeaders(v map[string]*string) *CreateAICoachTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAICoachTaskResponse) SetStatusCode(v int32) *CreateAICoachTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICoachTaskResponse) SetBody(v *CreateAICoachTaskResponseBody) *CreateAICoachTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAICoachTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
