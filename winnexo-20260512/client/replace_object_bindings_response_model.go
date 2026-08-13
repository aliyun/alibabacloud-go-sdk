// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceObjectBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceObjectBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceObjectBindingsResponseBody) *ReplaceObjectBindingsResponse
	GetBody() *ReplaceObjectBindingsResponseBody
}

type ReplaceObjectBindingsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceObjectBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceObjectBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsResponse) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceObjectBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceObjectBindingsResponse) GetBody() *ReplaceObjectBindingsResponseBody {
	return s.Body
}

func (s *ReplaceObjectBindingsResponse) SetHeaders(v map[string]*string) *ReplaceObjectBindingsResponse {
	s.Headers = v
	return s
}

func (s *ReplaceObjectBindingsResponse) SetStatusCode(v int32) *ReplaceObjectBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceObjectBindingsResponse) SetBody(v *ReplaceObjectBindingsResponseBody) *ReplaceObjectBindingsResponse {
	s.Body = v
	return s
}

func (s *ReplaceObjectBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
