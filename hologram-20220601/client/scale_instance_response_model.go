// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ScaleInstanceResponseBody) *ScaleInstanceResponse
	GetBody() *ScaleInstanceResponseBody
}

type ScaleInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleInstanceResponse) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleInstanceResponse) GetBody() *ScaleInstanceResponseBody {
	return s.Body
}

func (s *ScaleInstanceResponse) SetHeaders(v map[string]*string) *ScaleInstanceResponse {
	s.Headers = v
	return s
}

func (s *ScaleInstanceResponse) SetStatusCode(v int32) *ScaleInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleInstanceResponse) SetBody(v *ScaleInstanceResponseBody) *ScaleInstanceResponse {
	s.Body = v
	return s
}

func (s *ScaleInstanceResponse) Validate() error {
	return dara.Validate(s)
}
