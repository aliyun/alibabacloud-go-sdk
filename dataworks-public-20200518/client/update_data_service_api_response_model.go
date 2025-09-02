// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataServiceApiResponseBody) *UpdateDataServiceApiResponse
	GetBody() *UpdateDataServiceApiResponseBody
}

type UpdateDataServiceApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataServiceApiResponse) GetBody() *UpdateDataServiceApiResponseBody {
	return s.Body
}

func (s *UpdateDataServiceApiResponse) SetHeaders(v map[string]*string) *UpdateDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataServiceApiResponse) SetStatusCode(v int32) *UpdateDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataServiceApiResponse) SetBody(v *UpdateDataServiceApiResponseBody) *UpdateDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *UpdateDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
