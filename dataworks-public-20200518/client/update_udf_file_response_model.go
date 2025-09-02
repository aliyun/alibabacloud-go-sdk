// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUdfFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUdfFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUdfFileResponseBody) *UpdateUdfFileResponse
	GetBody() *UpdateUdfFileResponseBody
}

type UpdateUdfFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUdfFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateUdfFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUdfFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUdfFileResponse) GetBody() *UpdateUdfFileResponseBody {
	return s.Body
}

func (s *UpdateUdfFileResponse) SetHeaders(v map[string]*string) *UpdateUdfFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateUdfFileResponse) SetStatusCode(v int32) *UpdateUdfFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUdfFileResponse) SetBody(v *UpdateUdfFileResponseBody) *UpdateUdfFileResponse {
	s.Body = v
	return s
}

func (s *UpdateUdfFileResponse) Validate() error {
	return dara.Validate(s)
}
