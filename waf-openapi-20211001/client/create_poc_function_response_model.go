// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePocFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePocFunctionResponse
	GetStatusCode() *int32
	SetBody(v *CreatePocFunctionResponseBody) *CreatePocFunctionResponse
	GetBody() *CreatePocFunctionResponseBody
}

type CreatePocFunctionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePocFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePocFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePocFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreatePocFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePocFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePocFunctionResponse) GetBody() *CreatePocFunctionResponseBody {
	return s.Body
}

func (s *CreatePocFunctionResponse) SetHeaders(v map[string]*string) *CreatePocFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreatePocFunctionResponse) SetStatusCode(v int32) *CreatePocFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePocFunctionResponse) SetBody(v *CreatePocFunctionResponseBody) *CreatePocFunctionResponse {
	s.Body = v
	return s
}

func (s *CreatePocFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
