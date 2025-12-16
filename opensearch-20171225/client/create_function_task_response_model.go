// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFunctionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFunctionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFunctionTaskResponseBody) *CreateFunctionTaskResponse
	GetBody() *CreateFunctionTaskResponseBody
}

type CreateFunctionTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFunctionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFunctionTaskResponse) GetBody() *CreateFunctionTaskResponseBody {
	return s.Body
}

func (s *CreateFunctionTaskResponse) SetHeaders(v map[string]*string) *CreateFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionTaskResponse) SetStatusCode(v int32) *CreateFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionTaskResponse) SetBody(v *CreateFunctionTaskResponseBody) *CreateFunctionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFunctionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
