// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebPreciseAccessSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebPreciseAccessSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebPreciseAccessSwitchResponseBody) *ModifyWebPreciseAccessSwitchResponse
	GetBody() *ModifyWebPreciseAccessSwitchResponseBody
}

type ModifyWebPreciseAccessSwitchResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebPreciseAccessSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebPreciseAccessSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebPreciseAccessSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebPreciseAccessSwitchResponse) GetBody() *ModifyWebPreciseAccessSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebPreciseAccessSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebPreciseAccessSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebPreciseAccessSwitchResponse) SetStatusCode(v int32) *ModifyWebPreciseAccessSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchResponse) SetBody(v *ModifyWebPreciseAccessSwitchResponseBody) *ModifyWebPreciseAccessSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebPreciseAccessSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
