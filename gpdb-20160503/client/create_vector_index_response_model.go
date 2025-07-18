// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVectorIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVectorIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVectorIndexResponse
	GetStatusCode() *int32
	SetBody(v *CreateVectorIndexResponseBody) *CreateVectorIndexResponse
	GetBody() *CreateVectorIndexResponseBody
}

type CreateVectorIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVectorIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVectorIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVectorIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVectorIndexResponse) GetBody() *CreateVectorIndexResponseBody {
	return s.Body
}

func (s *CreateVectorIndexResponse) SetHeaders(v map[string]*string) *CreateVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateVectorIndexResponse) SetStatusCode(v int32) *CreateVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVectorIndexResponse) SetBody(v *CreateVectorIndexResponseBody) *CreateVectorIndexResponse {
	s.Body = v
	return s
}

func (s *CreateVectorIndexResponse) Validate() error {
	return dara.Validate(s)
}
