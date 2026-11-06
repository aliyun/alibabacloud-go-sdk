// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppInstanceGroupChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppInstanceGroupChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppInstanceGroupChargeTypeResponseBody) *ModifyAppInstanceGroupChargeTypeResponse
	GetBody() *ModifyAppInstanceGroupChargeTypeResponseBody
}

type ModifyAppInstanceGroupChargeTypeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInstanceGroupChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInstanceGroupChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) GetBody() *ModifyAppInstanceGroupChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceGroupChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) SetStatusCode(v int32) *ModifyAppInstanceGroupChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) SetBody(v *ModifyAppInstanceGroupChargeTypeResponseBody) *ModifyAppInstanceGroupChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
