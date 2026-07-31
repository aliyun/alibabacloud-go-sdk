// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSemanticViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSemanticViewResponse
	GetStatusCode() *int32
	SetBody(v *CreateSemanticViewResponseBody) *CreateSemanticViewResponse
	GetBody() *CreateSemanticViewResponseBody
}

type CreateSemanticViewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSemanticViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSemanticViewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticViewResponse) GoString() string {
	return s.String()
}

func (s *CreateSemanticViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSemanticViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSemanticViewResponse) GetBody() *CreateSemanticViewResponseBody {
	return s.Body
}

func (s *CreateSemanticViewResponse) SetHeaders(v map[string]*string) *CreateSemanticViewResponse {
	s.Headers = v
	return s
}

func (s *CreateSemanticViewResponse) SetStatusCode(v int32) *CreateSemanticViewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSemanticViewResponse) SetBody(v *CreateSemanticViewResponseBody) *CreateSemanticViewResponse {
	s.Body = v
	return s
}

func (s *CreateSemanticViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
