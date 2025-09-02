// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnabledExtensionsForProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnabledExtensionsForProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnabledExtensionsForProjectResponse
	GetStatusCode() *int32
	SetBody(v *ListEnabledExtensionsForProjectResponseBody) *ListEnabledExtensionsForProjectResponse
	GetBody() *ListEnabledExtensionsForProjectResponseBody
}

type ListEnabledExtensionsForProjectResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnabledExtensionsForProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnabledExtensionsForProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnabledExtensionsForProjectResponse) GoString() string {
	return s.String()
}

func (s *ListEnabledExtensionsForProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnabledExtensionsForProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnabledExtensionsForProjectResponse) GetBody() *ListEnabledExtensionsForProjectResponseBody {
	return s.Body
}

func (s *ListEnabledExtensionsForProjectResponse) SetHeaders(v map[string]*string) *ListEnabledExtensionsForProjectResponse {
	s.Headers = v
	return s
}

func (s *ListEnabledExtensionsForProjectResponse) SetStatusCode(v int32) *ListEnabledExtensionsForProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponse) SetBody(v *ListEnabledExtensionsForProjectResponseBody) *ListEnabledExtensionsForProjectResponse {
	s.Body = v
	return s
}

func (s *ListEnabledExtensionsForProjectResponse) Validate() error {
	return dara.Validate(s)
}
