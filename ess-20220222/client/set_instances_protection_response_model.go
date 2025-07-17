// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstancesProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetInstancesProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetInstancesProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetInstancesProtectionResponseBody) *SetInstancesProtectionResponse
	GetBody() *SetInstancesProtectionResponseBody
}

type SetInstancesProtectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstancesProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstancesProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetInstancesProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetInstancesProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetInstancesProtectionResponse) GetBody() *SetInstancesProtectionResponseBody {
	return s.Body
}

func (s *SetInstancesProtectionResponse) SetHeaders(v map[string]*string) *SetInstancesProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetInstancesProtectionResponse) SetStatusCode(v int32) *SetInstancesProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstancesProtectionResponse) SetBody(v *SetInstancesProtectionResponseBody) *SetInstancesProtectionResponse {
	s.Body = v
	return s
}

func (s *SetInstancesProtectionResponse) Validate() error {
	return dara.Validate(s)
}
