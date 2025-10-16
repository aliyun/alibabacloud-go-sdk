// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopChargeTypeResponseBody) *ModifyDesktopChargeTypeResponse
	GetBody() *ModifyDesktopChargeTypeResponseBody
}

type ModifyDesktopChargeTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopChargeTypeResponse) GetBody() *ModifyDesktopChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyDesktopChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyDesktopChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopChargeTypeResponse) SetStatusCode(v int32) *ModifyDesktopChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopChargeTypeResponse) SetBody(v *ModifyDesktopChargeTypeResponseBody) *ModifyDesktopChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
