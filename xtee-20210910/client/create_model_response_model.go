// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelResponseBody) *CreateModelResponse
	GetBody() *CreateModelResponseBody
}

type CreateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelResponse) GetBody() *CreateModelResponseBody {
	return s.Body
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetStatusCode(v int32) *CreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

func (s *CreateModelResponse) Validate() error {
	return dara.Validate(s)
}
