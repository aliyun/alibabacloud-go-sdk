// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQuotaAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQuotaAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CreateQuotaAlarmResponseBody) *CreateQuotaAlarmResponse
	GetBody() *CreateQuotaAlarmResponseBody
}

type CreateQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQuotaAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQuotaAlarmResponse) GetBody() *CreateQuotaAlarmResponseBody {
	return s.Body
}

func (s *CreateQuotaAlarmResponse) SetHeaders(v map[string]*string) *CreateQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaAlarmResponse) SetStatusCode(v int32) *CreateQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaAlarmResponse) SetBody(v *CreateQuotaAlarmResponseBody) *CreateQuotaAlarmResponse {
	s.Body = v
	return s
}

func (s *CreateQuotaAlarmResponse) Validate() error {
	return dara.Validate(s)
}
