// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateAICoachTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateAICoachTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateAICoachTaskResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateAICoachTaskResponseBody) *BatchCreateAICoachTaskResponse
	GetBody() *BatchCreateAICoachTaskResponseBody
}

type BatchCreateAICoachTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateAICoachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateAICoachTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateAICoachTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateAICoachTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateAICoachTaskResponse) GetBody() *BatchCreateAICoachTaskResponseBody {
	return s.Body
}

func (s *BatchCreateAICoachTaskResponse) SetHeaders(v map[string]*string) *BatchCreateAICoachTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateAICoachTaskResponse) SetStatusCode(v int32) *BatchCreateAICoachTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateAICoachTaskResponse) SetBody(v *BatchCreateAICoachTaskResponseBody) *BatchCreateAICoachTaskResponse {
	s.Body = v
	return s
}

func (s *BatchCreateAICoachTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
