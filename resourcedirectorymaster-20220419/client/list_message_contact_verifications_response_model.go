// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactVerificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageContactVerificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageContactVerificationsResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageContactVerificationsResponseBody) *ListMessageContactVerificationsResponse
	GetBody() *ListMessageContactVerificationsResponseBody
}

type ListMessageContactVerificationsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageContactVerificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageContactVerificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactVerificationsResponse) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageContactVerificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageContactVerificationsResponse) GetBody() *ListMessageContactVerificationsResponseBody {
	return s.Body
}

func (s *ListMessageContactVerificationsResponse) SetHeaders(v map[string]*string) *ListMessageContactVerificationsResponse {
	s.Headers = v
	return s
}

func (s *ListMessageContactVerificationsResponse) SetStatusCode(v int32) *ListMessageContactVerificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageContactVerificationsResponse) SetBody(v *ListMessageContactVerificationsResponseBody) *ListMessageContactVerificationsResponse {
	s.Body = v
	return s
}

func (s *ListMessageContactVerificationsResponse) Validate() error {
	return dara.Validate(s)
}
