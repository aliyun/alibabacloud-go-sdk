// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsDefaultSLRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenArmsDefaultSLRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenArmsDefaultSLRResponse
	GetStatusCode() *int32
	SetBody(v *OpenArmsDefaultSLRResponseBody) *OpenArmsDefaultSLRResponse
	GetBody() *OpenArmsDefaultSLRResponseBody
}

type OpenArmsDefaultSLRResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenArmsDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenArmsDefaultSLRResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenArmsDefaultSLRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenArmsDefaultSLRResponse) GetBody() *OpenArmsDefaultSLRResponseBody {
	return s.Body
}

func (s *OpenArmsDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenArmsDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetStatusCode(v int32) *OpenArmsDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetBody(v *OpenArmsDefaultSLRResponseBody) *OpenArmsDefaultSLRResponse {
	s.Body = v
	return s
}

func (s *OpenArmsDefaultSLRResponse) Validate() error {
	return dara.Validate(s)
}
