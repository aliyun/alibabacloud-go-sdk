// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserInformationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserInformationResponseBody) *UpdateUserInformationResponse
	GetBody() *UpdateUserInformationResponseBody
}

type UpdateUserInformationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserInformationResponse) GetBody() *UpdateUserInformationResponseBody {
	return s.Body
}

func (s *UpdateUserInformationResponse) SetHeaders(v map[string]*string) *UpdateUserInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserInformationResponse) SetStatusCode(v int32) *UpdateUserInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserInformationResponse) SetBody(v *UpdateUserInformationResponseBody) *UpdateUserInformationResponse {
	s.Body = v
	return s
}

func (s *UpdateUserInformationResponse) Validate() error {
	return dara.Validate(s)
}
