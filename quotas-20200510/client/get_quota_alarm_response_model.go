// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaAlarmResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaAlarmResponseBody) *GetQuotaAlarmResponse
	GetBody() *GetQuotaAlarmResponseBody
}

type GetQuotaAlarmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaAlarmResponse) GetBody() *GetQuotaAlarmResponseBody {
	return s.Body
}

func (s *GetQuotaAlarmResponse) SetHeaders(v map[string]*string) *GetQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaAlarmResponse) SetStatusCode(v int32) *GetQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaAlarmResponse) SetBody(v *GetQuotaAlarmResponseBody) *GetQuotaAlarmResponse {
	s.Body = v
	return s
}

func (s *GetQuotaAlarmResponse) Validate() error {
	return dara.Validate(s)
}
