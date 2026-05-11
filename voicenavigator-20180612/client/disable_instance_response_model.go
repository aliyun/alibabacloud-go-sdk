// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DisableInstanceResponseBody) *DisableInstanceResponse
	GetBody() *DisableInstanceResponseBody
}

type DisableInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceResponse) GoString() string {
	return s.String()
}

func (s *DisableInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInstanceResponse) GetBody() *DisableInstanceResponseBody {
	return s.Body
}

func (s *DisableInstanceResponse) SetHeaders(v map[string]*string) *DisableInstanceResponse {
	s.Headers = v
	return s
}

func (s *DisableInstanceResponse) SetStatusCode(v int32) *DisableInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstanceResponse) SetBody(v *DisableInstanceResponseBody) *DisableInstanceResponse {
	s.Body = v
	return s
}

func (s *DisableInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
