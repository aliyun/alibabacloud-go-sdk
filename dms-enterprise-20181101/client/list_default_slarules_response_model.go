// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDefaultSLARulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDefaultSLARulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDefaultSLARulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDefaultSLARulesResponseBody) *ListDefaultSLARulesResponse
	GetBody() *ListDefaultSLARulesResponseBody
}

type ListDefaultSLARulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDefaultSLARulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDefaultSLARulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultSLARulesResponse) GoString() string {
	return s.String()
}

func (s *ListDefaultSLARulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDefaultSLARulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDefaultSLARulesResponse) GetBody() *ListDefaultSLARulesResponseBody {
	return s.Body
}

func (s *ListDefaultSLARulesResponse) SetHeaders(v map[string]*string) *ListDefaultSLARulesResponse {
	s.Headers = v
	return s
}

func (s *ListDefaultSLARulesResponse) SetStatusCode(v int32) *ListDefaultSLARulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDefaultSLARulesResponse) SetBody(v *ListDefaultSLARulesResponseBody) *ListDefaultSLARulesResponse {
	s.Body = v
	return s
}

func (s *ListDefaultSLARulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
