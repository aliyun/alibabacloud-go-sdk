// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomCallTaggingResponseBody) *CreateCustomCallTaggingResponse
	GetBody() *CreateCustomCallTaggingResponseBody
}

type CreateCustomCallTaggingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomCallTaggingResponse) GetBody() *CreateCustomCallTaggingResponseBody {
	return s.Body
}

func (s *CreateCustomCallTaggingResponse) SetHeaders(v map[string]*string) *CreateCustomCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomCallTaggingResponse) SetStatusCode(v int32) *CreateCustomCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomCallTaggingResponse) SetBody(v *CreateCustomCallTaggingResponseBody) *CreateCustomCallTaggingResponse {
	s.Body = v
	return s
}

func (s *CreateCustomCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
