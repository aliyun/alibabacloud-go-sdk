// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpApiOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpApiOperationResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpApiOperationResponseBody) *CreateHttpApiOperationResponse
	GetBody() *CreateHttpApiOperationResponseBody
}

type CreateHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpApiOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpApiOperationResponse) GetBody() *CreateHttpApiOperationResponseBody {
	return s.Body
}

func (s *CreateHttpApiOperationResponse) SetHeaders(v map[string]*string) *CreateHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiOperationResponse) SetStatusCode(v int32) *CreateHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiOperationResponse) SetBody(v *CreateHttpApiOperationResponseBody) *CreateHttpApiOperationResponse {
	s.Body = v
	return s
}

func (s *CreateHttpApiOperationResponse) Validate() error {
	return dara.Validate(s)
}
