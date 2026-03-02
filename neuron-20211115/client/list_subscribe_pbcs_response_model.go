// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscribePbcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubscribePbcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubscribePbcsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubscribePbcsResponseBody) *ListSubscribePbcsResponse
	GetBody() *ListSubscribePbcsResponseBody
}

type ListSubscribePbcsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscribePbcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscribePbcsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubscribePbcsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscribePbcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubscribePbcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubscribePbcsResponse) GetBody() *ListSubscribePbcsResponseBody {
	return s.Body
}

func (s *ListSubscribePbcsResponse) SetHeaders(v map[string]*string) *ListSubscribePbcsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscribePbcsResponse) SetStatusCode(v int32) *ListSubscribePbcsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscribePbcsResponse) SetBody(v *ListSubscribePbcsResponseBody) *ListSubscribePbcsResponse {
	s.Body = v
	return s
}

func (s *ListSubscribePbcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
