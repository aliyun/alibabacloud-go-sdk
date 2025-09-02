// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeOwnerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeOwnerResponseBody) *UpdateNodeOwnerResponse
	GetBody() *UpdateNodeOwnerResponseBody
}

type UpdateNodeOwnerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOwnerResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeOwnerResponse) GetBody() *UpdateNodeOwnerResponseBody {
	return s.Body
}

func (s *UpdateNodeOwnerResponse) SetHeaders(v map[string]*string) *UpdateNodeOwnerResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeOwnerResponse) SetStatusCode(v int32) *UpdateNodeOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeOwnerResponse) SetBody(v *UpdateNodeOwnerResponseBody) *UpdateNodeOwnerResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeOwnerResponse) Validate() error {
	return dara.Validate(s)
}
