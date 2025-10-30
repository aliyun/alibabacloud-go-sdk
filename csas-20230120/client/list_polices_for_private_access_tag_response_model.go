// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicesForPrivateAccessTagResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicesForPrivateAccessTagResponseBody) *ListPolicesForPrivateAccessTagResponse
	GetBody() *ListPolicesForPrivateAccessTagResponseBody
}

type ListPolicesForPrivateAccessTagResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicesForPrivateAccessTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicesForPrivateAccessTagResponse) GetBody() *ListPolicesForPrivateAccessTagResponseBody {
	return s.Body
}

func (s *ListPolicesForPrivateAccessTagResponse) SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponse) SetStatusCode(v int32) *ListPolicesForPrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponse) SetBody(v *ListPolicesForPrivateAccessTagResponseBody) *ListPolicesForPrivateAccessTagResponse {
	s.Body = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
