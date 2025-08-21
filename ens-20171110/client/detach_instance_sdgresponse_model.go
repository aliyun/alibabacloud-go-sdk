// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *DetachInstanceSDGResponseBody) *DetachInstanceSDGResponse
	GetBody() *DetachInstanceSDGResponseBody
}

type DetachInstanceSDGResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachInstanceSDGResponse) GetBody() *DetachInstanceSDGResponseBody {
	return s.Body
}

func (s *DetachInstanceSDGResponse) SetHeaders(v map[string]*string) *DetachInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *DetachInstanceSDGResponse) SetStatusCode(v int32) *DetachInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachInstanceSDGResponse) SetBody(v *DetachInstanceSDGResponseBody) *DetachInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *DetachInstanceSDGResponse) Validate() error {
	return dara.Validate(s)
}
