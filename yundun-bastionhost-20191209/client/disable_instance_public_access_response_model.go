// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstancePublicAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInstancePublicAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInstancePublicAccessResponse
	GetStatusCode() *int32
	SetBody(v *DisableInstancePublicAccessResponseBody) *DisableInstancePublicAccessResponse
	GetBody() *DisableInstancePublicAccessResponseBody
}

type DisableInstancePublicAccessResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInstancePublicAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInstancePublicAccessResponse) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInstancePublicAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInstancePublicAccessResponse) GetBody() *DisableInstancePublicAccessResponseBody {
	return s.Body
}

func (s *DisableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *DisableInstancePublicAccessResponse {
	s.Headers = v
	return s
}

func (s *DisableInstancePublicAccessResponse) SetStatusCode(v int32) *DisableInstancePublicAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstancePublicAccessResponse) SetBody(v *DisableInstancePublicAccessResponseBody) *DisableInstancePublicAccessResponse {
	s.Body = v
	return s
}

func (s *DisableInstancePublicAccessResponse) Validate() error {
	return dara.Validate(s)
}
