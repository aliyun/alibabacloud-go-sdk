// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClaimAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClaimAlarmResponse
	GetStatusCode() *int32
	SetBody(v *ClaimAlarmResponseBody) *ClaimAlarmResponse
	GetBody() *ClaimAlarmResponseBody
}

type ClaimAlarmResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClaimAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClaimAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s ClaimAlarmResponse) GoString() string {
	return s.String()
}

func (s *ClaimAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClaimAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClaimAlarmResponse) GetBody() *ClaimAlarmResponseBody {
	return s.Body
}

func (s *ClaimAlarmResponse) SetHeaders(v map[string]*string) *ClaimAlarmResponse {
	s.Headers = v
	return s
}

func (s *ClaimAlarmResponse) SetStatusCode(v int32) *ClaimAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ClaimAlarmResponse) SetBody(v *ClaimAlarmResponseBody) *ClaimAlarmResponse {
	s.Body = v
	return s
}

func (s *ClaimAlarmResponse) Validate() error {
	return dara.Validate(s)
}
