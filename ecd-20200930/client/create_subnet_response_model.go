// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubnetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubnetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubnetResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubnetResponseBody) *CreateSubnetResponse
	GetBody() *CreateSubnetResponseBody
}

type CreateSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubnetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetResponse) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubnetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubnetResponse) GetBody() *CreateSubnetResponseBody {
	return s.Body
}

func (s *CreateSubnetResponse) SetHeaders(v map[string]*string) *CreateSubnetResponse {
	s.Headers = v
	return s
}

func (s *CreateSubnetResponse) SetStatusCode(v int32) *CreateSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubnetResponse) SetBody(v *CreateSubnetResponseBody) *CreateSubnetResponse {
	s.Body = v
	return s
}

func (s *CreateSubnetResponse) Validate() error {
	return dara.Validate(s)
}
