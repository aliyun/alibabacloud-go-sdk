// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesensitizeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DesensitizeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DesensitizeDataResponse
	GetStatusCode() *int32
	SetBody(v *DesensitizeDataResponseBody) *DesensitizeDataResponse
	GetBody() *DesensitizeDataResponseBody
}

type DesensitizeDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DesensitizeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DesensitizeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DesensitizeDataResponse) GoString() string {
	return s.String()
}

func (s *DesensitizeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DesensitizeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DesensitizeDataResponse) GetBody() *DesensitizeDataResponseBody {
	return s.Body
}

func (s *DesensitizeDataResponse) SetHeaders(v map[string]*string) *DesensitizeDataResponse {
	s.Headers = v
	return s
}

func (s *DesensitizeDataResponse) SetStatusCode(v int32) *DesensitizeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DesensitizeDataResponse) SetBody(v *DesensitizeDataResponseBody) *DesensitizeDataResponse {
	s.Body = v
	return s
}

func (s *DesensitizeDataResponse) Validate() error {
	return dara.Validate(s)
}
