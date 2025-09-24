// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ConvertChargeTypeResponseBody) *ConvertChargeTypeResponse
	GetBody() *ConvertChargeTypeResponseBody
}

type ConvertChargeTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertChargeTypeResponse) GetBody() *ConvertChargeTypeResponseBody {
	return s.Body
}

func (s *ConvertChargeTypeResponse) SetHeaders(v map[string]*string) *ConvertChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ConvertChargeTypeResponse) SetStatusCode(v int32) *ConvertChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertChargeTypeResponse) SetBody(v *ConvertChargeTypeResponseBody) *ConvertChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ConvertChargeTypeResponse) Validate() error {
	return dara.Validate(s)
}
