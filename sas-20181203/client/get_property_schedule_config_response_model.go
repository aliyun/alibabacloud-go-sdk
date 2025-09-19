// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertyScheduleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPropertyScheduleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPropertyScheduleConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPropertyScheduleConfigResponseBody) *GetPropertyScheduleConfigResponse
	GetBody() *GetPropertyScheduleConfigResponseBody
}

type GetPropertyScheduleConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPropertyScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPropertyScheduleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPropertyScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPropertyScheduleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPropertyScheduleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPropertyScheduleConfigResponse) GetBody() *GetPropertyScheduleConfigResponseBody {
	return s.Body
}

func (s *GetPropertyScheduleConfigResponse) SetHeaders(v map[string]*string) *GetPropertyScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPropertyScheduleConfigResponse) SetStatusCode(v int32) *GetPropertyScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPropertyScheduleConfigResponse) SetBody(v *GetPropertyScheduleConfigResponseBody) *GetPropertyScheduleConfigResponse {
	s.Body = v
	return s
}

func (s *GetPropertyScheduleConfigResponse) Validate() error {
	return dara.Validate(s)
}
