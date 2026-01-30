// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRenderingChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRenderingChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRenderingChargeTypeResponseBody) *ModifyRenderingChargeTypeResponse
	GetBody() *ModifyRenderingChargeTypeResponseBody
}

type ModifyRenderingChargeTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRenderingChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRenderingChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRenderingChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRenderingChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRenderingChargeTypeResponse) GetBody() *ModifyRenderingChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyRenderingChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyRenderingChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRenderingChargeTypeResponse) SetStatusCode(v int32) *ModifyRenderingChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRenderingChargeTypeResponse) SetBody(v *ModifyRenderingChargeTypeResponseBody) *ModifyRenderingChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyRenderingChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
