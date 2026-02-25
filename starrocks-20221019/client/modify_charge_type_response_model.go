// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyChargeTypeResponseBody) *ModifyChargeTypeResponse
	GetBody() *ModifyChargeTypeResponseBody
}

type ModifyChargeTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyChargeTypeResponse) GetBody() *ModifyChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyChargeTypeResponse) SetStatusCode(v int32) *ModifyChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyChargeTypeResponse) SetBody(v *ModifyChargeTypeResponseBody) *ModifyChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
