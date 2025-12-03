// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceAccessControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInstanceAccessControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInstanceAccessControlResponse
	GetStatusCode() *int32
	SetBody(v *DisableInstanceAccessControlResponseBody) *DisableInstanceAccessControlResponse
	GetBody() *DisableInstanceAccessControlResponseBody
}

type DisableInstanceAccessControlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstanceAccessControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInstanceAccessControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceAccessControlResponse) GoString() string {
	return s.String()
}

func (s *DisableInstanceAccessControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInstanceAccessControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInstanceAccessControlResponse) GetBody() *DisableInstanceAccessControlResponseBody {
	return s.Body
}

func (s *DisableInstanceAccessControlResponse) SetHeaders(v map[string]*string) *DisableInstanceAccessControlResponse {
	s.Headers = v
	return s
}

func (s *DisableInstanceAccessControlResponse) SetStatusCode(v int32) *DisableInstanceAccessControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstanceAccessControlResponse) SetBody(v *DisableInstanceAccessControlResponseBody) *DisableInstanceAccessControlResponse {
	s.Body = v
	return s
}

func (s *DisableInstanceAccessControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
