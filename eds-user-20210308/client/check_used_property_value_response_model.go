// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUsedPropertyValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUsedPropertyValueResponse
	GetStatusCode() *int32
	SetBody(v *CheckUsedPropertyValueResponseBody) *CheckUsedPropertyValueResponse
	GetBody() *CheckUsedPropertyValueResponseBody
}

type CheckUsedPropertyValueResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUsedPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUsedPropertyValueResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUsedPropertyValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUsedPropertyValueResponse) GetBody() *CheckUsedPropertyValueResponseBody {
	return s.Body
}

func (s *CheckUsedPropertyValueResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetStatusCode(v int32) *CheckUsedPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetBody(v *CheckUsedPropertyValueResponseBody) *CheckUsedPropertyValueResponse {
	s.Body = v
	return s
}

func (s *CheckUsedPropertyValueResponse) Validate() error {
	return dara.Validate(s)
}
