// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFieldResponse
	GetStatusCode() *int32
	SetBody(v *CreateFieldResponseBody) *CreateFieldResponse
	GetBody() *CreateFieldResponseBody
}

type CreateFieldResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFieldResponse) GoString() string {
	return s.String()
}

func (s *CreateFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFieldResponse) GetBody() *CreateFieldResponseBody {
	return s.Body
}

func (s *CreateFieldResponse) SetHeaders(v map[string]*string) *CreateFieldResponse {
	s.Headers = v
	return s
}

func (s *CreateFieldResponse) SetStatusCode(v int32) *CreateFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFieldResponse) SetBody(v *CreateFieldResponseBody) *CreateFieldResponse {
	s.Body = v
	return s
}

func (s *CreateFieldResponse) Validate() error {
	return dara.Validate(s)
}
