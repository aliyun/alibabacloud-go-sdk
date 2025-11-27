// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceChargeTypeResponseBody) *ModifyRCInstanceChargeTypeResponse
	GetBody() *ModifyRCInstanceChargeTypeResponseBody
}

type ModifyRCInstanceChargeTypeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceChargeTypeResponse) GetBody() *ModifyRCInstanceChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponse) SetStatusCode(v int32) *ModifyRCInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponse) SetBody(v *ModifyRCInstanceChargeTypeResponseBody) *ModifyRCInstanceChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
