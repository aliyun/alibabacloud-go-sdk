// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *GetAIAppJobResponseBody) *GetAIAppJobResponse
	GetBody() *GetAIAppJobResponseBody
}

type GetAIAppJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *GetAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIAppJobResponse) GetBody() *GetAIAppJobResponseBody {
	return s.Body
}

func (s *GetAIAppJobResponse) SetHeaders(v map[string]*string) *GetAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *GetAIAppJobResponse) SetStatusCode(v int32) *GetAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIAppJobResponse) SetBody(v *GetAIAppJobResponseBody) *GetAIAppJobResponse {
	s.Body = v
	return s
}

func (s *GetAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
