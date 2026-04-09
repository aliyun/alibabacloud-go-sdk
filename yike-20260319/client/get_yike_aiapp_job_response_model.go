// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeAIAppJobResponseBody) *GetYikeAIAppJobResponse
	GetBody() *GetYikeAIAppJobResponseBody
}

type GetYikeAIAppJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeAIAppJobResponse) GetBody() *GetYikeAIAppJobResponseBody {
	return s.Body
}

func (s *GetYikeAIAppJobResponse) SetHeaders(v map[string]*string) *GetYikeAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *GetYikeAIAppJobResponse) SetStatusCode(v int32) *GetYikeAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeAIAppJobResponse) SetBody(v *GetYikeAIAppJobResponseBody) *GetYikeAIAppJobResponse {
	s.Body = v
	return s
}

func (s *GetYikeAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
