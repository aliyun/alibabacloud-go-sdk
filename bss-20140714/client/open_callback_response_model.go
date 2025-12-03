// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenCallbackResponse
	GetStatusCode() *int32
	SetBody(v *OpenCallbackResponseBody) *OpenCallbackResponse
	GetBody() *OpenCallbackResponseBody
}

type OpenCallbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenCallbackResponse) GoString() string {
	return s.String()
}

func (s *OpenCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenCallbackResponse) GetBody() *OpenCallbackResponseBody {
	return s.Body
}

func (s *OpenCallbackResponse) SetHeaders(v map[string]*string) *OpenCallbackResponse {
	s.Headers = v
	return s
}

func (s *OpenCallbackResponse) SetStatusCode(v int32) *OpenCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCallbackResponse) SetBody(v *OpenCallbackResponseBody) *OpenCallbackResponse {
	s.Body = v
	return s
}

func (s *OpenCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
