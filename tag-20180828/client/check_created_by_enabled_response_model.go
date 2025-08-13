// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreatedByEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCreatedByEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCreatedByEnabledResponse
	GetStatusCode() *int32
	SetBody(v *CheckCreatedByEnabledResponseBody) *CheckCreatedByEnabledResponse
	GetBody() *CheckCreatedByEnabledResponseBody
}

type CheckCreatedByEnabledResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreatedByEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCreatedByEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCreatedByEnabledResponse) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCreatedByEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCreatedByEnabledResponse) GetBody() *CheckCreatedByEnabledResponseBody {
	return s.Body
}

func (s *CheckCreatedByEnabledResponse) SetHeaders(v map[string]*string) *CheckCreatedByEnabledResponse {
	s.Headers = v
	return s
}

func (s *CheckCreatedByEnabledResponse) SetStatusCode(v int32) *CheckCreatedByEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreatedByEnabledResponse) SetBody(v *CheckCreatedByEnabledResponseBody) *CheckCreatedByEnabledResponse {
	s.Body = v
	return s
}

func (s *CheckCreatedByEnabledResponse) Validate() error {
	return dara.Validate(s)
}
