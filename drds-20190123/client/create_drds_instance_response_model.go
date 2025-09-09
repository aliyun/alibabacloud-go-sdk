// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDrdsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDrdsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDrdsInstanceResponseBody) *CreateDrdsInstanceResponse
	GetBody() *CreateDrdsInstanceResponseBody
}

type CreateDrdsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDrdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDrdsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDrdsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDrdsInstanceResponse) GetBody() *CreateDrdsInstanceResponseBody {
	return s.Body
}

func (s *CreateDrdsInstanceResponse) SetHeaders(v map[string]*string) *CreateDrdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsInstanceResponse) SetStatusCode(v int32) *CreateDrdsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsInstanceResponse) SetBody(v *CreateDrdsInstanceResponseBody) *CreateDrdsInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDrdsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
