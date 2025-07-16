// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRangeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRangeResponseBody) *UpdateRangeResponse
	GetBody() *UpdateRangeResponseBody
}

type UpdateRangeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRangeResponse) GetBody() *UpdateRangeResponseBody {
	return s.Body
}

func (s *UpdateRangeResponse) SetHeaders(v map[string]*string) *UpdateRangeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRangeResponse) SetStatusCode(v int32) *UpdateRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRangeResponse) SetBody(v *UpdateRangeResponseBody) *UpdateRangeResponse {
	s.Body = v
	return s
}

func (s *UpdateRangeResponse) Validate() error {
	return dara.Validate(s)
}
