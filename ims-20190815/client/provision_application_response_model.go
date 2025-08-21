// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProvisionApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProvisionApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ProvisionApplicationResponseBody) *ProvisionApplicationResponse
	GetBody() *ProvisionApplicationResponseBody
}

type ProvisionApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProvisionApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProvisionApplicationResponse) GetBody() *ProvisionApplicationResponseBody {
	return s.Body
}

func (s *ProvisionApplicationResponse) SetHeaders(v map[string]*string) *ProvisionApplicationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionApplicationResponse) SetStatusCode(v int32) *ProvisionApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionApplicationResponse) SetBody(v *ProvisionApplicationResponseBody) *ProvisionApplicationResponse {
	s.Body = v
	return s
}

func (s *ProvisionApplicationResponse) Validate() error {
	return dara.Validate(s)
}
