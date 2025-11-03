// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DestroyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DestroyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DestroyInstanceResponseBody) *DestroyInstanceResponse
	GetBody() *DestroyInstanceResponseBody
}

type DestroyInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DestroyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DestroyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DestroyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DestroyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DestroyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DestroyInstanceResponse) GetBody() *DestroyInstanceResponseBody {
	return s.Body
}

func (s *DestroyInstanceResponse) SetHeaders(v map[string]*string) *DestroyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DestroyInstanceResponse) SetStatusCode(v int32) *DestroyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DestroyInstanceResponse) SetBody(v *DestroyInstanceResponseBody) *DestroyInstanceResponse {
	s.Body = v
	return s
}

func (s *DestroyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
