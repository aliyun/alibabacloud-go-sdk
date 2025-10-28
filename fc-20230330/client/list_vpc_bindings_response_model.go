// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcBindingsOutput) *ListVpcBindingsResponse
	GetBody() *ListVpcBindingsOutput
}

type ListVpcBindingsResponse struct {
	Headers    map[string]*string     `json:"headers" xml:"headers"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcBindingsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcBindingsResponse) GetBody() *ListVpcBindingsOutput {
	return s.Body
}

func (s *ListVpcBindingsResponse) SetHeaders(v map[string]*string) *ListVpcBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcBindingsResponse) SetStatusCode(v int32) *ListVpcBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcBindingsResponse) SetBody(v *ListVpcBindingsOutput) *ListVpcBindingsResponse {
	s.Body = v
	return s
}

func (s *ListVpcBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
