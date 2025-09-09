// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDrdsDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDrdsDBResponse
	GetStatusCode() *int32
	SetBody(v *CreateDrdsDBResponseBody) *CreateDrdsDBResponse
	GetBody() *CreateDrdsDBResponseBody
}

type CreateDrdsDBResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDrdsDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDrdsDBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDrdsDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDrdsDBResponse) GetBody() *CreateDrdsDBResponseBody {
	return s.Body
}

func (s *CreateDrdsDBResponse) SetHeaders(v map[string]*string) *CreateDrdsDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDrdsDBResponse) SetStatusCode(v int32) *CreateDrdsDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDrdsDBResponse) SetBody(v *CreateDrdsDBResponseBody) *CreateDrdsDBResponse {
	s.Body = v
	return s
}

func (s *CreateDrdsDBResponse) Validate() error {
	return dara.Validate(s)
}
