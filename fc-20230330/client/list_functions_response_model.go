// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionsResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionsOutput) *ListFunctionsResponse
	GetBody() *ListFunctionsOutput
}

type ListFunctionsResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionsResponse) GetBody() *ListFunctionsOutput {
	return s.Body
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsOutput) *ListFunctionsResponse {
	s.Body = v
	return s
}

func (s *ListFunctionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
