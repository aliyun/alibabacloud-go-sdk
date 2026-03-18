// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckYikeAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckYikeAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckYikeAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckYikeAIAppJobResponseBody) *PrecheckYikeAIAppJobResponse
	GetBody() *PrecheckYikeAIAppJobResponseBody
}

type PrecheckYikeAIAppJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckYikeAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckYikeAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckYikeAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *PrecheckYikeAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckYikeAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckYikeAIAppJobResponse) GetBody() *PrecheckYikeAIAppJobResponseBody {
	return s.Body
}

func (s *PrecheckYikeAIAppJobResponse) SetHeaders(v map[string]*string) *PrecheckYikeAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *PrecheckYikeAIAppJobResponse) SetStatusCode(v int32) *PrecheckYikeAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckYikeAIAppJobResponse) SetBody(v *PrecheckYikeAIAppJobResponseBody) *PrecheckYikeAIAppJobResponse {
	s.Body = v
	return s
}

func (s *PrecheckYikeAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
