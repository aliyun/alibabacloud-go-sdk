// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsureOrderPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsureOrderPayResponse
	GetStatusCode() *int32
	SetBody(v *InsureOrderPayResponseBody) *InsureOrderPayResponse
	GetBody() *InsureOrderPayResponseBody
}

type InsureOrderPayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsureOrderPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsureOrderPayResponse) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderPayResponse) GoString() string {
	return s.String()
}

func (s *InsureOrderPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsureOrderPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsureOrderPayResponse) GetBody() *InsureOrderPayResponseBody {
	return s.Body
}

func (s *InsureOrderPayResponse) SetHeaders(v map[string]*string) *InsureOrderPayResponse {
	s.Headers = v
	return s
}

func (s *InsureOrderPayResponse) SetStatusCode(v int32) *InsureOrderPayResponse {
	s.StatusCode = &v
	return s
}

func (s *InsureOrderPayResponse) SetBody(v *InsureOrderPayResponseBody) *InsureOrderPayResponse {
	s.Body = v
	return s
}

func (s *InsureOrderPayResponse) Validate() error {
	return dara.Validate(s)
}
