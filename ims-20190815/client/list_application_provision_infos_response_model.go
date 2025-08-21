// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationProvisionInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationProvisionInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationProvisionInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationProvisionInfosResponseBody) *ListApplicationProvisionInfosResponse
	GetBody() *ListApplicationProvisionInfosResponseBody
}

type ListApplicationProvisionInfosResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationProvisionInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationProvisionInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationProvisionInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationProvisionInfosResponse) GetBody() *ListApplicationProvisionInfosResponseBody {
	return s.Body
}

func (s *ListApplicationProvisionInfosResponse) SetHeaders(v map[string]*string) *ListApplicationProvisionInfosResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationProvisionInfosResponse) SetStatusCode(v int32) *ListApplicationProvisionInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationProvisionInfosResponse) SetBody(v *ListApplicationProvisionInfosResponseBody) *ListApplicationProvisionInfosResponse {
	s.Body = v
	return s
}

func (s *ListApplicationProvisionInfosResponse) Validate() error {
	return dara.Validate(s)
}
