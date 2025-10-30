// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountFactoryBaselinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountFactoryBaselinesResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountFactoryBaselinesResponseBody) *ListAccountFactoryBaselinesResponse
	GetBody() *ListAccountFactoryBaselinesResponseBody
}

type ListAccountFactoryBaselinesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountFactoryBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountFactoryBaselinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountFactoryBaselinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountFactoryBaselinesResponse) GetBody() *ListAccountFactoryBaselinesResponseBody {
	return s.Body
}

func (s *ListAccountFactoryBaselinesResponse) SetHeaders(v map[string]*string) *ListAccountFactoryBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListAccountFactoryBaselinesResponse) SetStatusCode(v int32) *ListAccountFactoryBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponse) SetBody(v *ListAccountFactoryBaselinesResponseBody) *ListAccountFactoryBaselinesResponse {
	s.Body = v
	return s
}

func (s *ListAccountFactoryBaselinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
