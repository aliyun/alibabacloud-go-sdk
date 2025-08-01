// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnsServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnsServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListAnsServicesResponseBody) *ListAnsServicesResponse
	GetBody() *ListAnsServicesResponseBody
}

type ListAnsServicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnsServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnsServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServicesResponse) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnsServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnsServicesResponse) GetBody() *ListAnsServicesResponseBody {
	return s.Body
}

func (s *ListAnsServicesResponse) SetHeaders(v map[string]*string) *ListAnsServicesResponse {
	s.Headers = v
	return s
}

func (s *ListAnsServicesResponse) SetStatusCode(v int32) *ListAnsServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnsServicesResponse) SetBody(v *ListAnsServicesResponseBody) *ListAnsServicesResponse {
	s.Body = v
	return s
}

func (s *ListAnsServicesResponse) Validate() error {
	return dara.Validate(s)
}
