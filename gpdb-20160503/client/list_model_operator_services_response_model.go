// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelOperatorServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelOperatorServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelOperatorServicesResponseBody) *ListModelOperatorServicesResponse
	GetBody() *ListModelOperatorServicesResponseBody
}

type ListModelOperatorServicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelOperatorServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelOperatorServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorServicesResponse) GoString() string {
	return s.String()
}

func (s *ListModelOperatorServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelOperatorServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelOperatorServicesResponse) GetBody() *ListModelOperatorServicesResponseBody {
	return s.Body
}

func (s *ListModelOperatorServicesResponse) SetHeaders(v map[string]*string) *ListModelOperatorServicesResponse {
	s.Headers = v
	return s
}

func (s *ListModelOperatorServicesResponse) SetStatusCode(v int32) *ListModelOperatorServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelOperatorServicesResponse) SetBody(v *ListModelOperatorServicesResponseBody) *ListModelOperatorServicesResponse {
	s.Body = v
	return s
}

func (s *ListModelOperatorServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
