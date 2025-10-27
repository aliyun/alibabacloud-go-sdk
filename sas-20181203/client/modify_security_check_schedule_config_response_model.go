// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityCheckScheduleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityCheckScheduleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityCheckScheduleConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityCheckScheduleConfigResponseBody) *ModifySecurityCheckScheduleConfigResponse
	GetBody() *ModifySecurityCheckScheduleConfigResponseBody
}

type ModifySecurityCheckScheduleConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityCheckScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityCheckScheduleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityCheckScheduleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityCheckScheduleConfigResponse) GetBody() *ModifySecurityCheckScheduleConfigResponseBody {
	return s.Body
}

func (s *ModifySecurityCheckScheduleConfigResponse) SetHeaders(v map[string]*string) *ModifySecurityCheckScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityCheckScheduleConfigResponse) SetStatusCode(v int32) *ModifySecurityCheckScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigResponse) SetBody(v *ModifySecurityCheckScheduleConfigResponseBody) *ModifySecurityCheckScheduleConfigResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityCheckScheduleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
