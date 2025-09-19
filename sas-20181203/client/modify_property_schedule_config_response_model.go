// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPropertyScheduleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPropertyScheduleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPropertyScheduleConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPropertyScheduleConfigResponseBody) *ModifyPropertyScheduleConfigResponse
	GetBody() *ModifyPropertyScheduleConfigResponseBody
}

type ModifyPropertyScheduleConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPropertyScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPropertyScheduleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPropertyScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyPropertyScheduleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPropertyScheduleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPropertyScheduleConfigResponse) GetBody() *ModifyPropertyScheduleConfigResponseBody {
	return s.Body
}

func (s *ModifyPropertyScheduleConfigResponse) SetHeaders(v map[string]*string) *ModifyPropertyScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyPropertyScheduleConfigResponse) SetStatusCode(v int32) *ModifyPropertyScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPropertyScheduleConfigResponse) SetBody(v *ModifyPropertyScheduleConfigResponseBody) *ModifyPropertyScheduleConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyPropertyScheduleConfigResponse) Validate() error {
	return dara.Validate(s)
}
