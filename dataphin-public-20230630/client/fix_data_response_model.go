// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FixDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FixDataResponse
	GetStatusCode() *int32
	SetBody(v *FixDataResponseBody) *FixDataResponse
	GetBody() *FixDataResponseBody
}

type FixDataResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FixDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FixDataResponse) String() string {
	return dara.Prettify(s)
}

func (s FixDataResponse) GoString() string {
	return s.String()
}

func (s *FixDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FixDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FixDataResponse) GetBody() *FixDataResponseBody {
	return s.Body
}

func (s *FixDataResponse) SetHeaders(v map[string]*string) *FixDataResponse {
	s.Headers = v
	return s
}

func (s *FixDataResponse) SetStatusCode(v int32) *FixDataResponse {
	s.StatusCode = &v
	return s
}

func (s *FixDataResponse) SetBody(v *FixDataResponseBody) *FixDataResponse {
	s.Body = v
	return s
}

func (s *FixDataResponse) Validate() error {
	return dara.Validate(s)
}
