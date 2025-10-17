// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCategoryCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCategoryCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCategoryCallbackResponse
	GetStatusCode() *int32
	SetBody(v *StopCategoryCallbackResponseBody) *StopCategoryCallbackResponse
	GetBody() *StopCategoryCallbackResponseBody
}

type StopCategoryCallbackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCategoryCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCategoryCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCategoryCallbackResponse) GoString() string {
	return s.String()
}

func (s *StopCategoryCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCategoryCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCategoryCallbackResponse) GetBody() *StopCategoryCallbackResponseBody {
	return s.Body
}

func (s *StopCategoryCallbackResponse) SetHeaders(v map[string]*string) *StopCategoryCallbackResponse {
	s.Headers = v
	return s
}

func (s *StopCategoryCallbackResponse) SetStatusCode(v int32) *StopCategoryCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCategoryCallbackResponse) SetBody(v *StopCategoryCallbackResponseBody) *StopCategoryCallbackResponse {
	s.Body = v
	return s
}

func (s *StopCategoryCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
