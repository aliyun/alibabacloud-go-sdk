// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigItemsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigItemsResponseBody) *UpdateConfigItemsResponse
	GetBody() *UpdateConfigItemsResponseBody
}

type UpdateConfigItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigItemsResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigItemsResponse) GetBody() *UpdateConfigItemsResponseBody {
	return s.Body
}

func (s *UpdateConfigItemsResponse) SetHeaders(v map[string]*string) *UpdateConfigItemsResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigItemsResponse) SetStatusCode(v int32) *UpdateConfigItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigItemsResponse) SetBody(v *UpdateConfigItemsResponseBody) *UpdateConfigItemsResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigItemsResponse) Validate() error {
	return dara.Validate(s)
}
