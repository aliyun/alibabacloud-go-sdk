// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCategoryCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCategoryCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCategoryCallbackResponse
	GetStatusCode() *int32
	SetBody(v *StartCategoryCallbackResponseBody) *StartCategoryCallbackResponse
	GetBody() *StartCategoryCallbackResponseBody
}

type StartCategoryCallbackResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCategoryCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCategoryCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCategoryCallbackResponse) GoString() string {
	return s.String()
}

func (s *StartCategoryCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCategoryCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCategoryCallbackResponse) GetBody() *StartCategoryCallbackResponseBody {
	return s.Body
}

func (s *StartCategoryCallbackResponse) SetHeaders(v map[string]*string) *StartCategoryCallbackResponse {
	s.Headers = v
	return s
}

func (s *StartCategoryCallbackResponse) SetStatusCode(v int32) *StartCategoryCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCategoryCallbackResponse) SetBody(v *StartCategoryCallbackResponseBody) *StartCategoryCallbackResponse {
	s.Body = v
	return s
}

func (s *StartCategoryCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
