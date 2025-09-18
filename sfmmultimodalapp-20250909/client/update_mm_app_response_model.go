// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppResponseBody) *UpdateMmAppResponse
	GetBody() *UpdateMmAppResponseBody
}

type UpdateMmAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppResponse) GetBody() *UpdateMmAppResponseBody {
	return s.Body
}

func (s *UpdateMmAppResponse) SetHeaders(v map[string]*string) *UpdateMmAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppResponse) SetStatusCode(v int32) *UpdateMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppResponse) SetBody(v *UpdateMmAppResponseBody) *UpdateMmAppResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppResponse) Validate() error {
	return dara.Validate(s)
}
