// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVirtualMFAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableVirtualMFAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableVirtualMFAResponse
	GetStatusCode() *int32
	SetBody(v *DisableVirtualMFAResponseBody) *DisableVirtualMFAResponse
	GetBody() *DisableVirtualMFAResponseBody
}

type DisableVirtualMFAResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVirtualMFAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVirtualMFAResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableVirtualMFAResponse) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableVirtualMFAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableVirtualMFAResponse) GetBody() *DisableVirtualMFAResponseBody {
	return s.Body
}

func (s *DisableVirtualMFAResponse) SetHeaders(v map[string]*string) *DisableVirtualMFAResponse {
	s.Headers = v
	return s
}

func (s *DisableVirtualMFAResponse) SetStatusCode(v int32) *DisableVirtualMFAResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVirtualMFAResponse) SetBody(v *DisableVirtualMFAResponseBody) *DisableVirtualMFAResponse {
	s.Body = v
	return s
}

func (s *DisableVirtualMFAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
