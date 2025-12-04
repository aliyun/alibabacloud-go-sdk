// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubnetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubnetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubnetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSubnetResponseBody) *UpdateSubnetResponse
	GetBody() *UpdateSubnetResponseBody
}

type UpdateSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubnetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubnetResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubnetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubnetResponse) GetBody() *UpdateSubnetResponseBody {
	return s.Body
}

func (s *UpdateSubnetResponse) SetHeaders(v map[string]*string) *UpdateSubnetResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubnetResponse) SetStatusCode(v int32) *UpdateSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubnetResponse) SetBody(v *UpdateSubnetResponseBody) *UpdateSubnetResponse {
	s.Body = v
	return s
}

func (s *UpdateSubnetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
