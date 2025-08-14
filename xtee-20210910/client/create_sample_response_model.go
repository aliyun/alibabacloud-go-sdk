// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSampleResponse
	GetStatusCode() *int32
	SetBody(v *CreateSampleResponseBody) *CreateSampleResponse
	GetBody() *CreateSampleResponseBody
}

type CreateSampleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSampleResponse) GetBody() *CreateSampleResponseBody {
	return s.Body
}

func (s *CreateSampleResponse) SetHeaders(v map[string]*string) *CreateSampleResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleResponse) SetStatusCode(v int32) *CreateSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleResponse) SetBody(v *CreateSampleResponseBody) *CreateSampleResponse {
	s.Body = v
	return s
}

func (s *CreateSampleResponse) Validate() error {
	return dara.Validate(s)
}
