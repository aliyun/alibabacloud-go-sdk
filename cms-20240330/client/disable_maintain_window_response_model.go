// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMaintainWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableMaintainWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableMaintainWindowResponse
	GetStatusCode() *int32
	SetBody(v *DisableMaintainWindowResponseBody) *DisableMaintainWindowResponse
	GetBody() *DisableMaintainWindowResponseBody
}

type DisableMaintainWindowResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableMaintainWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableMaintainWindowResponse) GoString() string {
	return s.String()
}

func (s *DisableMaintainWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableMaintainWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableMaintainWindowResponse) GetBody() *DisableMaintainWindowResponseBody {
	return s.Body
}

func (s *DisableMaintainWindowResponse) SetHeaders(v map[string]*string) *DisableMaintainWindowResponse {
	s.Headers = v
	return s
}

func (s *DisableMaintainWindowResponse) SetStatusCode(v int32) *DisableMaintainWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMaintainWindowResponse) SetBody(v *DisableMaintainWindowResponseBody) *DisableMaintainWindowResponse {
	s.Body = v
	return s
}

func (s *DisableMaintainWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
