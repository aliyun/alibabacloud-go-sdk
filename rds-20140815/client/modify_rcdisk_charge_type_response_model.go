// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCDiskChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCDiskChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCDiskChargeTypeResponseBody) *ModifyRCDiskChargeTypeResponse
	GetBody() *ModifyRCDiskChargeTypeResponseBody
}

type ModifyRCDiskChargeTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCDiskChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCDiskChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCDiskChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCDiskChargeTypeResponse) GetBody() *ModifyRCDiskChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyRCDiskChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyRCDiskChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCDiskChargeTypeResponse) SetStatusCode(v int32) *ModifyRCDiskChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCDiskChargeTypeResponse) SetBody(v *ModifyRCDiskChargeTypeResponseBody) *ModifyRCDiskChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCDiskChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
