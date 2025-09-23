// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelServicesResponseBody) *ListModelServicesResponse
	GetBody() *ListModelServicesResponseBody
}

type ListModelServicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesResponse) GoString() string {
	return s.String()
}

func (s *ListModelServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelServicesResponse) GetBody() *ListModelServicesResponseBody {
	return s.Body
}

func (s *ListModelServicesResponse) SetHeaders(v map[string]*string) *ListModelServicesResponse {
	s.Headers = v
	return s
}

func (s *ListModelServicesResponse) SetStatusCode(v int32) *ListModelServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelServicesResponse) SetBody(v *ListModelServicesResponseBody) *ListModelServicesResponse {
	s.Body = v
	return s
}

func (s *ListModelServicesResponse) Validate() error {
	return dara.Validate(s)
}
