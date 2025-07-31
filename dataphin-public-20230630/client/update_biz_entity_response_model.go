// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBizEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBizEntityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBizEntityResponseBody) *UpdateBizEntityResponse
	GetBody() *UpdateBizEntityResponseBody
}

type UpdateBizEntityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBizEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBizEntityResponse) GetBody() *UpdateBizEntityResponseBody {
	return s.Body
}

func (s *UpdateBizEntityResponse) SetHeaders(v map[string]*string) *UpdateBizEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizEntityResponse) SetStatusCode(v int32) *UpdateBizEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizEntityResponse) SetBody(v *UpdateBizEntityResponseBody) *UpdateBizEntityResponse {
	s.Body = v
	return s
}

func (s *UpdateBizEntityResponse) Validate() error {
	return dara.Validate(s)
}
