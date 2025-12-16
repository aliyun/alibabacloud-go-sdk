// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFunctionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFunctionResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateFunctionResourceResponseBody) *CreateFunctionResourceResponse
	GetBody() *CreateFunctionResourceResponseBody
}

type CreateFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFunctionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFunctionResourceResponse) GetBody() *CreateFunctionResourceResponseBody {
	return s.Body
}

func (s *CreateFunctionResourceResponse) SetHeaders(v map[string]*string) *CreateFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResourceResponse) SetStatusCode(v int32) *CreateFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResourceResponse) SetBody(v *CreateFunctionResourceResponseBody) *CreateFunctionResourceResponse {
	s.Body = v
	return s
}

func (s *CreateFunctionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
