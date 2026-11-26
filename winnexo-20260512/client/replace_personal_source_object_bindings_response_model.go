// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplacePersonalSourceObjectBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplacePersonalSourceObjectBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplacePersonalSourceObjectBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ReplacePersonalSourceObjectBindingsResponseBody) *ReplacePersonalSourceObjectBindingsResponse
	GetBody() *ReplacePersonalSourceObjectBindingsResponseBody
}

type ReplacePersonalSourceObjectBindingsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplacePersonalSourceObjectBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplacePersonalSourceObjectBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsResponse) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplacePersonalSourceObjectBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplacePersonalSourceObjectBindingsResponse) GetBody() *ReplacePersonalSourceObjectBindingsResponseBody {
	return s.Body
}

func (s *ReplacePersonalSourceObjectBindingsResponse) SetHeaders(v map[string]*string) *ReplacePersonalSourceObjectBindingsResponse {
	s.Headers = v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponse) SetStatusCode(v int32) *ReplacePersonalSourceObjectBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponse) SetBody(v *ReplacePersonalSourceObjectBindingsResponseBody) *ReplacePersonalSourceObjectBindingsResponse {
	s.Body = v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
